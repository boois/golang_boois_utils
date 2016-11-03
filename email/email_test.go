package email

import (
	"testing"
	"fmt"
)

func TestSendToMail(t *testing.T) {

	user := "log@toalaska.net"
	password := ""
	host := "smtp.mxhichina.com:25"
	to := "1726950105@qq.com"

	subject := "中虽dddsdf国"

	body := `
		<html>
		<body>
		<h3>
		呵呵
		</h3>
		使用Golang发送邮件1111使用Golang发送邮件1111使用Golang发送邮件1111使用Golang发送邮件1111使用Golang发送邮件1111
		<img src="https://www.baidu.com/img/baidu_jgylogo3.gif">
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendEmail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
