package common

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

const (
	//BASE64字符表,不要有重复
	base64Table        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	hashFunctionHeader = ""
	hashFunctionFooter = ""
)

var coder = base64.NewEncoding(base64Table)

func Md5(value string, salt string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	var result []byte
	if strings.EqualFold(salt, "") {
		result = hash.Sum(nil)
	} else {
		result = hash.Sum([]byte(salt))
	}
	return hex.EncodeToString(result)
}
func SHA1(value string, salt string) string {

	t := sha1.New()
	t.Write([]byte(value))
	var result []byte
	if strings.EqualFold(salt, "") {
		result = t.Sum(nil)
	} else {
		result = t.Sum([]byte(salt))
	}
	return hex.EncodeToString(result)
}

/**
 * base64加密
 */
func Base64Encode(str string) string {
	var src []byte = []byte(hashFunctionHeader + str + hashFunctionFooter)
	return string([]byte(coder.EncodeToString(src)))
}

/**
 * base64解密
 */
func Base64Decode(str string) (string, error) {
	var src []byte = []byte(str)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}
