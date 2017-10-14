package error

// 阿里云API公共错误码, 参考 https://help.aliyun.com/document_detail/53438.html
const (
	OK                        = "OK"                                         //请求成功
	ProductNotFound           = "InvalidProduct.NotFound"                    //通过域名找不到对应的产品。
	ApiNotFound               = "InvalidApi.NotFound"                        //找不到指定的 API
	ProtocolNeedSsl           = "InvalidProtocol.NeedSsl"                    //只支持 HTTPS 协议
	AccessKeyIdNotFound       = "InvalidAccessKeyId.NotFound"                //AccessKeyId 找不到
	AccessKeyIdInactive       = "InvalidAccessKeyId.Inactive"                //AccessKeyId 被禁用
	TimeStampFormat           = "InvalidTimeStamp.Format"                    //时间戳格式不对
	TimeStampExpired          = "InvalidTimeStamp.Expired"                   //用户时间和服务器时间不在 15 分钟内
	ParameterFormat           = "InvalidParameter.Format"                    //返回值格式不正确（Format 不支持）
	ParameterAccept           = "InvalidParameter.Accept"                    //返回值格式不正确（Accept 不支持）
	TokenExpired              = "InvalidSecurityToken.Expired"               //SecurityToken 过期
	TokenMismatchAccessKey    = "InvalidSecurityToken.MismatchWithAccessKey" //SecurityToken 与 AccessKey 不匹配
	TokenMalformed            = "InvalidSecurityToken.Malformed"             //SecurityToken 错误
	InvalidParameterName      = "Invalid{ParameterName}"                     //参数值或名称校验不通过
	InvalidParameter          = "InvalidParameter"                           //API参数值校验不通过
	InvalidSignatureMethod    = "InvalidSignatureMethod"                     //签名方法不支持
	IncompleteSignature       = "IncompleteSignature"                        //签名不匹配
	RegionNotFound            = "InvalidRegion.NotFound"                     //找不到请求对应的 Region
	InternalError             = "InternalError"                              //内部错误
	ServiceUnavailable        = "ServiceUnavailable"                         //服务暂时不可用（底层服务不可用）
	SignatureDoesNotMatch     = "SignatureDoesNotMatch"                      //签名不匹配
	SignatureNonceUsed        = "SignatureNonceUsed"                         //SignatureNonce 重复
	MissingSecurityToken      = "MissingSecurityToken"                       //缺少 SecurityToken
	MissingParameter          = "MissingParameter"                           //必填参数没有填
	MissingParameterName      = "Missing{ParameterName}"                     //必填参数没有填
	ContentLengthDoesNotMatch = "ContentLengthDoesNotMatch"                  //指定的 content-length 与 body 长度不匹配
	ContentMD5NotMatched      = "ContentMD5NotMatched"                       //MD5 校验不通过
	UnsupportedHTTPMethod     = "UnsupportedHTTPMethod"                      //HTTP 请求方法不支持
	AccessKeyDisabled         = "Forbidden.AccessKeyDisabled"                //AccessKeyId 被禁用
	Throttling                = "Throttling"                                 //您这个时段的流量已经超限
	ThrottlingUser            = "Throttling.User"                            //您这个时段的流量已经超限
	ThrottlingApi             = "Throttling.Api"                             //您这个时段的流量已经超限
	UnknownError              = "UnknownError"                               //未知错误
)
