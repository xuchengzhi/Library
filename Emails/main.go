package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
	// "strings"
)

type EmailStr struct {
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	ToUers string
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCUers []string
}

var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message

func InitEmail(els EmailStr) {
	// toers := []string{}

	serverHost = els.ServerHost
	serverPort = els.ServerPort
	fromEmail = els.FromEmail
	fromPasswd = els.FromPasswd

	m = gomail.NewMessage()
	if len(els.ToUers) == 0 {
		return
	}

	// for _, tmp := range strings.Split(ep.Toers, ",") {
	// 	toers = append(toers, strings.TrimSpace(tmp))
	// }

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", els.ToUers)
	fmt.Println(els.CCUers)
	// m.SetHeader("Cc", els.CCUers)
	//抄送列表
	// if len(els.CCUers) != 0 {
	// 	for _, tmp := range strings.Split(els.CCUers, ",") {
	// 		toers = append(toers, strings.TrimSpace(tmp))
	// 	}
	// 	fmt.Println(toers)
	// 	// m.SetHeader("Cc", toers...)
	// }

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", fromEmail, "测试管理员")
}

func SendEmail(subject, body string) {
	// 主题
	m.SetHeader("Subject", subject)

	// 正文
	m.SetBody("text/html", body)
	fmt.Println(serverHost)
	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}

func main() {
	var subject = "this is my emails"
	var body = "我也不知道该写啥"

	// fromEmail = "xuchengzhi1987@yeah.net"
	// fromPasswd = "xcz258521"
	// serverPort = 25
	// serverHost = "smtp.yeah.net"
	var els EmailStr
	els.ServerHost = "smtp.yeah.net"

	els.ServerPort = 25

	els.FromEmail = "xuchengzhi1987@yeah.net"

	els.FromPasswd = "xcz258521"

	els.ToUers = "2654676540@qq.com"

	els.CCUers = []string{"xudear@yeah.net", "foundertest@yeah.net"}
	InitEmail(els)
	SendEmail(subject, body)
}
