package xiaomiio

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/icepie/micloud"
	"github.com/tidwall/gjson"
)

const (
	APP_DEVICE_ID       = "3C861A5820190429"
	SDK_VER             = "3.9"
	APP_UA              = "APP/com.xiaomi.mihome APPV/6.0.103 iosPassportSDK/3.9.0 iOS/14.4 miHSTS"
	MINA_UA             = "MiHome/6.0.103 (com.xiaomi.mihome; build:6.0.103.1; iOS 14.4.0) Alamofire/6.0.103 MICO/iOSApp/appStore/6.0.103"
	MIIO_UA             = "iOS-14.4-6.0.103-iPhone12,3--D7744744F7AF32F0544445285880DD63E47D9BE9-8816080-84A3F44E137B71AE-iPhone"
	MIIO_SID            = "xiaomiio"
	Default_Country     = "CN"
	SERVICE_LOGIN_AUTH2 = "https://account.xiaomi.com/pass/serviceLoginAuth2"
	SERVICE_LOGIN       = "https://account.xiaomi.com/pass/serviceLogin"
)

//
type XiaoMiio struct {
	Client    *resty.Client
	Password  string
	Username  string
	Country   string
	CUserId   string
	UserId    int64
	PassToken string
	// Location      string
	Nonce         string
	SecurityToken string
}

// NewClient creates a new client
func NewXiaoMiio(username string, password string) *XiaoMiio {

	client := resty.New()

	// client.SetDebug(true)

	client.SetHeaders(map[string]string{
		"User-Agent": MIIO_UA,
		"Accept":     "*/*",
		"Connection": "keep-alive",
	})

	client.SetCookies([]*http.Cookie{
		{
			Name:  "sdkVersion",
			Value: SDK_VER,
		},
		{
			Name:  "deviceId",
			Value: micloud.GenRandomDeviceID(),
		},
	})

	return &XiaoMiio{
		Client:   client,
		Password: password,
		Username: username,
		Country:  Default_Country,
	}
}

func (xm *XiaoMiio) loginStep1() (jsonStr string, err error) {

	resp, err := xm.Client.R().SetQueryParams(
		map[string]string{
			"sid":   MIIO_SID,
			"_json": "true",
		},
	).Get(SERVICE_LOGIN)

	if err != nil {
		return
	}

	jsonStr = resp.String()

	return
}

func (xm *XiaoMiio) loginStep2(authSign string) (locationUrl string, err error) {

	if xm.Username == "" || xm.Password == "" {
		err = errors.New("username or password is empty")
		return
	}

	resp, err := xm.Client.R().SetQueryParams(
		map[string]string{
			"user":     xm.Username,
			"hash":     strings.ToUpper(micloud.MD5(xm.Password)),
			"_sign":    gjson.Get(authSign, "_sign").String(),
			"qs":       gjson.Get(authSign, "qs").String(),
			"callback": gjson.Get(authSign, "callback").String(),
			"sid":      MIIO_SID,
			"_json":    "true",
		},
	).Post(SERVICE_LOGIN_AUTH2)

	if err != nil {
		return
	}

	jsonStr := resp.String()

	if gjson.Get(jsonStr, "code").Int() != 0 {
		err = errors.New(gjson.Get(jsonStr, "description").String())
		return
	}

	xm.SecurityToken = gjson.Get(jsonStr, "ssecurity").String()
	xm.PassToken = gjson.Get(jsonStr, "psecurity").String()
	xm.UserId = gjson.Get(jsonStr, "userId").Int()
	xm.Nonce = gjson.Get(jsonStr, "nonce").String()

	xm.CUserId = gjson.Get(jsonStr, "cUserId").String()

	locationUrl = gjson.Get(jsonStr, "location").String()

	return

}

func (xm *XiaoMiio) loginStep3(locationUrl string) (err error) {

	resp, err := xm.Client.R().SetQueryParams(
		map[string]string{
			"sid":   MIIO_SID,
			"_json": "true",
		},
	).Get(locationUrl)

	if err != nil {
		return
	}

	log.Println(resp.Cookies())

	return

}

// 是否登录
func (xm *XiaoMiio) IsLogin() bool {

	jsonStr, err := xm.loginStep1()
	if err != nil {
		// log.Println(err)
		return false
	}

	return gjson.Get(jsonStr, "code").Int() == 0
}

func (xm *XiaoMiio) Login() (err error) {

	jsonStr, err := xm.loginStep1()
	if err != nil {
		log.Println(err)
		return
	}

	locationUrl, err := xm.loginStep2(jsonStr)
	if err != nil {
		log.Println(err)
		return
	}

	err = xm.loginStep3(locationUrl)
	if err != nil {
		log.Println(err)
		return
	}

	// log.Println(locationUrl)

	return

}
