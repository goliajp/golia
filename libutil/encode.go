package libutil

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

func Md5(data string) string {
	md5Obj := md5.New()
	md5Obj.Write([]byte(data))
	md5Data := md5Obj.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func Hmac(key, data string) string {
	hmacObj := hmac.New(md5.New, []byte(key))
	hmacObj.Write([]byte(data))
	return hex.EncodeToString(hmacObj.Sum([]byte("")))
}

func Sha1(data string) string {
	sha1Obj := sha1.New()
	sha1Obj.Write([]byte(data))
	return hex.EncodeToString(sha1Obj.Sum([]byte("")))
}

func UrlEncode(data string) string {
	return url.QueryEscape(data)
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) string {
	decode, _ := base64.StdEncoding.DecodeString(data)
	return string(decode)
}
