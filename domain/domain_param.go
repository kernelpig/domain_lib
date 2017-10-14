package domain

import (
	"github.com/satori/go.uuid"

	"wangqingang/sms_lib/common"
	"wangqingang/sms_lib/pb"
)

// 设置系统参数
func FillPublicParam(params map[string]string) {
	params["SignatureMethod"] = common.SignMethod
	params["SignatureNonce"] = uuid.NewV4().String()
	params["AccessKeyId"] = AccessKeyId
	params["SignatureVersion"] = common.SignVersion
	params["Timestamp"] = common.CurrentDateTimeFormat()
	params["Format"] = common.SmsSendResponseFormat
}

// 设置业务参数
func FillBusinessParam(params map[string]string, smsReq *pb.SendSmsRequest) {
	params["Action"] = common.SmsSendAction
	params["Version"] = common.SmsSendVersion
	params["RegionId"] = common.SmsSendRegionId
	params["PhoneNumbers"] = smsReq.PhoneNumbers
	params["SignName"] = smsReq.SignName
	params["TemplateParam"] = smsReq.TemplateParam
	params["TemplateCode"] = smsReq.TemplateCode
}

// 设置业务参数
func FillParam(params map[string]string, smsReq *pb.SendSmsRequest) {
	FillPublicParam(params)
	FillBusinessParam(params, smsReq)

	// 去除签名关键字Key
	if _, ok := params["Signature"]; ok {
		delete(params, "Signature")
	}
}
