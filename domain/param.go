package domain

import (
	"fmt"
	"strconv"

	"github.com/satori/go.uuid"

	"wangqingang/domain_lib/common"
	"wangqingang/domain_lib/pb"
)

// 业务请求处理常量
const (
	reqKeyAction                 = "Action"
	reqKeySubOrder               = "SubOrderParam"
	reqKeySubOrderPeriodFmt      = reqKeySubOrder + ".%d.Period"
	reqKeySubOrderActionFmt      = reqKeySubOrder + ".%d.Action"
	reqKeySubOrderRelatedNameFmt = reqKeySubOrder + ".%d.RelatedName"
	reqKeySubOrderTemplateIDFmt  = reqKeySubOrder + ".%d.DomainTemplateID"
)

// 业务API接口参数常量，
const (
	createOrderAction = "CreateOrder"
)

// 设置系统参数
func p2rPublicParam(params map[string]string) {
	params[common.ReqKeySignatureMethod] = common.SignMethod
	params[common.ReqKeySignatureNonce] = uuid.NewV4().String()
	params[common.ReqKeyAccessKeyId] = AccessKeyId
	params[common.ReqKeySignatureVersion] = common.SignVersion
	params[common.ReqKeyTimestamp] = common.CurrentDateTimeFormat()
	params[common.ReqKeyFormat] = common.ResponseFormat
	params[common.ReqKeyVersion] = ReqAPIVersion
}

// 转换子订单参数
func p2rSubOrderParam(params map[string]string, req []*pb.SubOrderParam) {
	for i := 0; i < len(req); i++ {
		a := fmt.Sprintf(reqKeySubOrderActionFmt, i+1)
		p := fmt.Sprintf(reqKeySubOrderPeriodFmt, i+1)
		t := fmt.Sprintf(reqKeySubOrderTemplateIDFmt, i+1)
		rn := fmt.Sprintf(reqKeySubOrderRelatedNameFmt, i+1)
		params[a] = req[i].Action
		params[rn] = req[i].RelatedName
		params[t] = req[i].DomainTemplateID
		params[p] = strconv.Itoa(int(req[i].Period))
	}
}

// 转换创建订单参数
func p2rCreateOrderReq(params map[string]string, req *pb.CreateOrderRequest) {
	params[reqKeyAction] = createOrderAction
	p2rSubOrderParam(params, req.SubOrderParam)
}

// 设置业务参数
func fillParam(params map[string]string, req *pb.CreateOrderRequest) {
	p2rPublicParam(params)
	p2rCreateOrderReq(params, req)

	// 去除签名关键字Key
	if _, ok := params[common.ReqKeySignature]; ok {
		delete(params, common.ReqKeySignature)
	}
}
