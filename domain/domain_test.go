package domain

import (
	"testing"
	"wangqingang/sms_lib/pb"
)

const (
	testAccessId      = "YOUR_ACCESS_ID"
	testAccessSecret  = "YOUR_ACCESS_SECRET"
	testPhoneNumbers  = "18653183000"
	testSignName      = "大脸"
	testTemplateCode  = "SMS_89765057"
	testTemplateParam = "{\"code\": \"123456\"}"
)

func TestCreateSmsSendUrl(t *testing.T) {
	InitAccess(testAccessId, testAccessSecret)

	request := pb.SendSmsRequest{
		PhoneNumbers:  testPhoneNumbers,
		SignName:      testSignName,
		TemplateCode:  testTemplateCode,
		TemplateParam: testTemplateParam,
	}
	url := CreateSmsSendUrl(&request)
	if len(url) == 0 {
		t.Error("create sms send url failed.")
	}
	t.Log("your send sms url is: ", url)
}

func TestCreateSmsSendUrlWithAccess(t *testing.T) {
	request := pb.SendSmsRequest{
		PhoneNumbers:  testPhoneNumbers,
		SignName:      testSignName,
		TemplateCode:  testTemplateCode,
		TemplateParam: testTemplateParam,
	}
	url := CreateSmsSendUrlWithAccess(testAccessId, testAccessSecret, &request)
	if len(url) == 0 {
		t.Error("create sms send url failed.")
	}
	t.Log("your send sms url is: ", url)
}
