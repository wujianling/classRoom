package sms

import "fmt"

func SendSms(phone, method, content string) (err error) {
	fmt.Printf("成功发送短信 手机号：%s ，类型：%s ， 信息：%s 。 \n", phone, method, content)
	return nil
}
