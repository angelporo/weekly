package mail

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/liyuan/weekly/html"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	FileNameDir string
	Title       string
}

// send function
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ":")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func (M *Mail) Send() error {
	// monthMap := map[string]string{
	//	"January":   "01",
	//	"February":  "02",
	//	"March":     "03",
	//	"April":     "04",
	//	"May":       "05",
	//	"June":      "06",
	//	"July":      "07",
	//	"August":    "08",
	//	"September": "09",
	//	"October":   "10",
	//	"November":  "11",
	//	"December":  "12",
	// }
	// t := time.Now()
	// y, month, day := t.Date()

	m := gomail.NewMessage()

	m.SetAddressHeader("From", "liyuan@bettem.com" /*"发件人地址"*/, "李渊") // 发件人

	m.SetHeader("To",
		m.FormatAddress("b@bettem.com", "收件人")) // 收件人
	m.SetHeader("Cc",
		m.FormatAddress("man@bettem.com", "收件人")) //抄送
	m.SetHeader("Cc",
		m.FormatAddress("mal@bettem.com", "收件人")) //抄送
	m.SetHeader("Cc",
		m.FormatAddress("jijm@bettem.com", "收件人")) //抄送

	m.SetHeader("Subject", M.Title) // 主题

	m.SetBody("text/html", html.Html) // 可以放html..还有其他的
	// m.SetBody(body) // 正文
	m.Attach(M.FileNameDir) //添加附件

	d := gomail.NewPlainDialer("smtp.qiye.163.com", 25, "liyuan@bettem.com", "1pLvzZ9HNvcmWSRn") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return err
	}

	os.Remove(M.FileNameDir)
	log.Println("发送成功! 并删除excel文件成功")
	return nil
}
