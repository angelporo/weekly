package mail

import (
	"log"
	"os"

	"github.com/liyuan/weekly/html"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	FileNameDir string   // 文件绝对路径
	Title       string   // 标题
	UserName    string   // 用户名
	Point       int      // 开放端口
	Auth        string   // 作者 -> 邮件别名
	PassWord    string   // 密码 -> 这里需要 smtp 授权码
	SendTo      string   // 发送目标
	CopyTo      []string // 抄送目标
	Host        string   // 发送目标域
}

type Cc struct {
	Email string
	Alias string
	Type  string // Cc 抄送目标 To 发送目标
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

	var auth string
	if M.Auth == "" {
		auth = "李渊"
	} else {
		auth = M.Auth
	}
	m.SetAddressHeader("From", M.UserName /*"发件人地址"*/, auth) // 发件人

	m.SetHeader("To",
		m.FormatAddress(M.SendTo, "收件人")) // 收件人

	m.SetHeader("Cc",
		M.CopyTo...) //抄送

	m.SetHeader("Subject", M.Title) // 主题

	m.SetBody("text/html", html.Html) // 可以放html..还有其他的
	m.Attach(M.FileNameDir)           //添加附件

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewPlainDialer(M.Host, M.Point, M.UserName, M.PassWord)
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败: 如有疑问请联系李渊", err)
		return err
	}

	os.Remove(M.FileNameDir)
	log.Println("发送成功! 并删除 excel 文件成功!")
	return nil
}
