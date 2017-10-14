package common

const (
	SignMethod       = "HMAC-SHA1"
	SignVersion      = "1.0"
	SignStringFormat = "GET&%s&%s"
	SignSecretFormat = "%s&"
	utcTimeFormat    = "2006-01-02T15:04:05Z"

	SmsSendAction         = "SendSms"
	SmsSendVersion        = "2017-05-25"
	SmsSendRegionId       = "cn-hangzhou"
	SmsSendResponseFormat = "JSON"
	SmsSendUrlFormat      = "http://dysmsapi.aliyuncs.com/?Signature=%s&%s"
)
