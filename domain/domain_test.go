package domain

import (
	"testing"
	"wangqingang/domain_lib/pb"
)

const (
	testAccessId      = "YOUR_ACCESS_ID"
	testAccessSecret  = "YOUR_ACCESS_SECRET"
	testAction       = createOrderAction
	testDomain1      = "yx.xin"
	testTemplateID   = "3834855"
)

func TestCreateSmsSendUrl(t *testing.T) {
	InitAccess(testAccessId, testAccessSecret)

	subOrder := &pb.SubOrderParam{
		RelatedName:      testDomain1,
		Period:           12,
		DomainTemplateID: testTemplateID,
		Action:           SubOrderActionActivate,
	}
	request := pb.CreateOrderRequest{
		Action:        testAction,
		SubOrderParam: []*pb.SubOrderParam{subOrder},
	}
	url := CreateSmsSendUrl(&request)
	if len(url) == 0 {
		t.Error("create domain create order url failed.")
	}
	t.Log("your domain create order url is: ", url)
}

//func TestCreateSmsSendUrlWithAccess(t *testing.T) {
//	request := pb.SendSmsRequest{
//		PhoneNumbers:  testPhoneNumbers,
//		SignName:      testSignName,
//		TemplateCode:  testTemplateCode,
//		TemplateParam: testTemplateParam,
//	}
//	url := CreateSmsSendUrlWithAccess(testAccessId, testAccessSecret, &request)
//	if len(url) == 0 {
//		t.Error("create sms send url failed.")
//	}
//	t.Log("your send sms url is: ", url)
//}
