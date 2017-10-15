package domain

import (
	"fmt"
	"sort"
	"strings"

	"wangqingang/domain_lib/common"
	"wangqingang/domain_lib/pb"
)

// 业务请求参数
const (
	ReqURI                 = "http://domain.aliyuncs.com/?Signature=%s&%s"
	ReqAPIVersion          = "2016-05-11"
	SubOrderActionActivate = "activate"
	SubOrderActionRenew    = "renew"
)

var AccessKeyId string
var AccessSecret string

// 初始化阿里云访问参数
func InitAccess(keyId, secret string) {
	AccessKeyId = keyId
	AccessSecret = secret
}

// 创建域名发送链接
func CreateOrderWithAccess(accessKeyId, accessSecret string, req *pb.CreateOrderRequest) string {
	InitAccess(accessKeyId, accessSecret)
	return CreateOrder(req)
}

// 创建域名发送链接
func CreateOrder(req *pb.CreateOrderRequest) string {
	params := make(map[string]string)

	// 设置请求参数
	fillParam(params, req)

	// 参数Key排序, 获取排序后的key
	sortedKeys := make([]string, 0)
	for key, _ := range params {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	// 构造待签名的字符串, 只对key/value 特殊url编码
	sortQueryStringSet := make([]string, 0)
	for _, key := range sortedKeys {
		encodingKey := common.SpecialUrlEncode(key)
		encodingValue := common.SpecialUrlEncode(params[key])
		encodingQuery := fmt.Sprintf("%s=%s", encodingKey, encodingValue)
		sortQueryStringSet = append(sortQueryStringSet, encodingQuery)
	}
	sortedQueryString := strings.Join(sortQueryStringSet, "&")

	// 生成签名字符串
	encodingChar := common.SpecialUrlEncode("/")
	encodingSortedQuery := common.SpecialUrlEncode(sortedQueryString)
	signingSecret := fmt.Sprintf(common.SignSecretFormat, AccessSecret)
	signingContent := fmt.Sprintf(common.SignStringFormat, encodingChar, encodingSortedQuery)
	signedString := common.Sign(signingSecret, signingContent)

	// 签名最后也要做特殊URL编码
	signature := common.SpecialUrlEncode(signedString)

	// 返回合法的请求
	return fmt.Sprintf(ReqURI, signature, sortedQueryString)
}
