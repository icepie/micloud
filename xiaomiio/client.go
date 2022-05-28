package xiaomiio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/icepie/micloud"
	"github.com/tidwall/gjson"
)

const (
	// APP_DEVICE_ID       = "3C861A5820190429"
	SDK_VER             = "3.9"
	APP_UA              = "APP/com.xiaomi.mihome APPV/6.0.103 iosPassportSDK/3.9.0 iOS/14.4 miHSTS"
	MINA_UA             = "MiHome/6.0.103 (com.xiaomi.mihome; build:6.0.103.1; iOS 14.4.0) Alamofire/6.0.103 MICO/iOSApp/appStore/6.0.103"
	MIIO_UA             = "iOS-14.4-6.0.103-iPhone12,3--D7744744F7AF32F0544445285880DD63E47D9BE9-8816080-84A3F44E137B71AE-iPhone"
	MIIO_SID            = "xiaomiio"
	MIIO_BASE_API       = "https://api.io.mi.com/app"
	MIIO_I18N_API       = "https://%s.api.io.mi.com/app/i18n"
	DEFAULT_COUNTRY     = "cn"
	SERVICE_LOGIN_AUTH2 = "https://account.xiaomi.com/pass/serviceLoginAuth2"
	SERVICE_LOGIN       = "https://account.xiaomi.com/pass/serviceLogin"

	//  Error define
	ERROR_NOT_LOGGED = "you are not logged in"
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
		Country:  DEFAULT_COUNTRY,
	}
}

// Set the country
func (xm *XiaoMiio) SetCountry(country string) *XiaoMiio {
	xm.Country = strings.ToLower(country)
	return xm
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

	if gjson.Get(jsonStr, "result").String() != "ok" {
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

	xm.Client.SetCookies(resp.Cookies())

	return

}

// 是否登录
func (xm *XiaoMiio) IsLogged() bool {

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
		return
	}

	locationUrl, err := xm.loginStep2(jsonStr)
	if err != nil {
		return
	}

	err = xm.loginStep3(locationUrl)
	if err != nil {
		return
	}

	// log.Println(locationUrl)

	return

}

func (xm *XiaoMiio) getRequestUrl() string {
	if xm.Country == DEFAULT_COUNTRY {
		return MIIO_BASE_API
	}

	return fmt.Sprintf(MIIO_I18N_API, xm.Country)
}

// data: json string
func (xm *XiaoMiio) Request(path string, data string) (ret string, err error) {

	if !xm.IsLogged() {
		err = errors.New(ERROR_NOT_LOGGED)
		return
	}

	url := xm.getRequestUrl() + path

	nonce, err := micloud.GenNonce()
	if err != nil {
		return
	}

	signedNonce, err := micloud.GenSignedNonce(xm.SecurityToken, nonce)
	if err != nil {
		return
	}

	signature, err := micloud.GenSignature(path, signedNonce, nonce, data)
	if err != nil {
		return
	}

	// var ret RespRet

	resp, err := xm.Client.R().
		SetQueryParams(
			map[string]string{
				"signature": signature,
				"_nonce":    nonce,
				"data":      data,
			},
		).
		SetHeaders(map[string]string{
			"x-xiaomi-protocal-flag-cli": "PROTOCAL-HTTP2",
		}).
		Get(url)

	ret = resp.String()

	return

}

func (xm *XiaoMiio) GetDevicesByRaw(req GetDevicesReq) (devices []Device, err error) {

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return
	}

	resp, err := xm.Request("/home/device_list", string(jsonBytes))
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(gjson.Get(resp, "result.list").String()), &devices)

	return
}

func (xm *XiaoMiio) GetDevices() (devices []Device, err error) {
	return xm.GetDevicesByRaw(GetDevicesReq{
		GetVirtualModel: false,
		GetHuamiDevices: 0,
	})
}

func (xm *XiaoMiio) GetProps(params ...PropParam) (ret []PropRet, err error) {
	paramsReq := PropParamsReq{
		Params: params,
	}

	jsonBytes, err := json.Marshal(paramsReq)
	if err != nil {
		return
	}

	resp, err := xm.Request("/miotspec/prop/get", string(jsonBytes))
	if err != nil {
		return
	}

	if gjson.Get(resp, "code").Int() != 0 {
		err = errors.New(gjson.Get(resp, "message").String())
		return
	}

	json.Unmarshal([]byte(gjson.Get(resp, "result").String()), &ret)

	return

}

func (xm *XiaoMiio) SetProps(params ...PropParam) (ret []PropRet, err error) {

	paramsReq := PropParamsReq{
		Params: params,
	}

	jsonBytes, err := json.Marshal(paramsReq)
	if err != nil {
		return
	}

	resp, err := xm.Request("/miotspec/prop/set", string(jsonBytes))
	if err != nil {
		return
	}

	if gjson.Get(resp, "code").Int() != 0 {
		err = errors.New(gjson.Get(resp, "message").String())
		return
	}

	json.Unmarshal([]byte(gjson.Get(resp, "result").String()), &ret)

	return

}

func (xm *XiaoMiio) DoAction(param ActionParam) (ret ActionRet, err error) {

	if param.In == nil {
		param.In = []interface{}{}
	}

	req := ActionParamReq{
		Param: param,
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return
	}

	resp, err := xm.Request("/miotspec/action", string(jsonBytes))
	if err != nil {
		return
	}

	if gjson.Get(resp, "code").Int() != 0 {
		err = errors.New(gjson.Get(resp, "message").String())
		return
	}

	json.Unmarshal([]byte(gjson.Get(resp, "result").String()), &ret)

	return

}

// func (xm *XiaoMiio) GetUserDeviceDataByRaw(did string) (err error) {

// 	// resp, err := xm.Request("/miotspec/action", string(jsonBytes))
// 	// if err != nil {
// 	// 	return
// 	// }

// }
