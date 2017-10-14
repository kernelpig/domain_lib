package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"strings"
	"time"
)

func CurrentDateTimeFormat() string {
	return time.Now().UTC().Format(utcTimeFormat)
}

func SpecialUrlEncode(s string) string {
	encodedString := url.QueryEscape(s)
	encodedString = strings.Replace(encodedString, "+", "%20", -1)
	encodedString = strings.Replace(encodedString, "*", "%2A", -1)
	encodedString = strings.Replace(encodedString, "%7E", "~", -1)
	return encodedString
}

func Sign(accessSecret, stringToSign string) string {
	h := hmac.New(sha1.New, []byte(accessSecret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
