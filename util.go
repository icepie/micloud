package micloud

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"time"
)

// 随机生成16位字符串的设备ID
func GenRandomDeviceID() string {
	return "3C861A5820190419"
}

// md5 加密
func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func GenNonce() string {
	tb := make([]byte, 4)
	binary.LittleEndian.PutUint32(tb, uint32(time.Now().Unix()/60))

	nb := make([]byte, 8)
	_, _ = rand.Read(nb)

	return base64.StdEncoding.EncodeToString(append(nb, tb...))
}

// Nonce signed with ssecret
func GetSignedNonce(ssecret string, nonce string) (signedNonce string, err error) {
	h := sha256.New()

	ssecretDecode, err := base64.StdEncoding.DecodeString(ssecret)

	if err != nil {
		return
	}

	nonceDecode, err := base64.StdEncoding.DecodeString(nonce)
	if err != nil {
		return
	}

	h.Write(ssecretDecode)
	h.Write(nonceDecode)

	signedNonce = base64.StdEncoding.EncodeToString(h.Sum(nil))

	return
}

// Request signature based on url, signed_nonce, nonce and data

func GenSignature(path string, signedNonce string, nonce string, data string) (signature string, err error) {
	sign := strings.Join([]string{path, signedNonce, nonce, "data=" + data}, "&")

	signedNonceDecode, err := base64.StdEncoding.DecodeString(signedNonce)
	if err != nil {
		return
	}

	h := hmac.New(sha256.New, signedNonceDecode)

	h.Write([]byte(sign))

	signature = base64.StdEncoding.EncodeToString(h.Sum(nil))

	return
}
