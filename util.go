package micloud

import (
	"crypto/md5"
	"encoding/hex"
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
