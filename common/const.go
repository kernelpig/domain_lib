package common

// 公共请求参数KEY
const (
	ReqKeySignatureMethod  = "SignatureMethod"
	ReqKeySignatureNonce   = "SignatureNonce"
	ReqKeyAccessKeyId      = "AccessKeyId"
	ReqKeySignatureVersion = "SignatureVersion"
	ReqKeyTimestamp        = "Timestamp"
	ReqKeyFormat           = "Format"
	ReqKeyVersion          = "Version"
	ReqKeySignature        = "Signature"
)

// 公共请求参数值格式化
const (
	SignMethod       = "HMAC-SHA1"
	SignVersion      = "1.0"
	SignStringFormat = "GET&%s&%s"
	SignSecretFormat = "%s&"
	utcTimeFormat    = "2006-01-02T15:04:05Z"
	ResponseFormat   = "JSON"
)
