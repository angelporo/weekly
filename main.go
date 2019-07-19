package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/liyuan/weekly/subjoin"
	"github.com/urfave/cli"
	"gopkg.in/gomail.v2"
)

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

const text = `<div id="qm_con_body"><div id="mailContentContainer" class="qmbox qm_con_body_content qqmail_webmail_only" style=""><div style="line-height:1.7;color:#000000;font-size:14px;font-family:Arial"><br><br><br><br><br><div style="position:relative;zoom:1">--<br><div><span style=" font-size: 13px ; ;; ">李渊 | 软件开发部</span></div><div><span style=" font-size: 13px ; ;; "><span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="176-3669-9981">176-3669-9981</span> | liyuan<a href="mailto:guojia@bettem.com" src="mailto:guojia@bettem.com" rel="noopener" target="_blank">@bettem.com</a></span></div><div><br></div><div><span style=" font-size: 13px ; ;; "><b>山西百得科技开发股份有限公司</b></span></div><div><span style=" font-size: 13px ; ;; ">股票代码 | 839289</span></div><div><span style="  ;; font-size: 13px; ">公司地址：山西 | 太原 南中环街401号数码港A座二层</span></div><div><span style="  ;; font-size: 13px; ">公司网址：<a>www.bettem.com</a> </span></div><div><span style="  ;; font-size: 13px; ">公司电话：<span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="0351-7033691">0351-7033691</span></span></div><div style="clear:both"></div></div><div id="dvLetterAngle"> </div></div><style type="text/css">.qmbox style, .qmbox script, .qmbox head, .qmbox link, .qmbox meta {display: none !important;}</style></div></div>`

const body = `
    <html>
    <body>
    <h3>
    ` + text + `
    </h3>
    </body>
    </html>
    `

func main() {
	app := cli.NewApp()
	app.Name = "weekly"
	app.Usage = "use: weekly <starttime> <endtime> <weekly content> "

	app.Action = func(c *cli.Context) error {
		Exc := subjoin.Excel{
			Start:   c.Args().Get(0),
			End:     c.Args().Get(1),
			Content: c.Args().Get(2),
		}
		if err := Exc.NewExcel(); err != nil {
			return err
		}
		// filename := Exc.GetFileDir()

		// send(filename)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func send(filename string) {
	monthMap := map[string]string{
		"January":   "01",
		"February":  "02",
		"March":     "03",
		"April":     "04",
		"May":       "05",
		"June":      "06",
		"July":      "07",
		"August":    "08",
		"September": "09",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}
	t := time.Now()
	y, month, day := t.Date()

	// FIXME:邮件主题 时间需要修改
	title := fmt.Sprint("软件部李渊工作周报", y, monthMap[month.String()], day)

	m := gomail.NewMessage()

	m.SetAddressHeader("From", "liyuan@bettem.com" /*"发件人地址"*/, "李渊") // 发件人

	m.SetHeader("To",
		m.FormatAddress("940079461@qq.com", "收件人")) // 收件人
	// m.SetHeader("Cc",
	//	m.FormatAddress("ruanb@bettem.com", "收件人")) //抄送
	// m.SetHeader("Cc",
	//	m.FormatAddress("mal@bettem.com", "收件人")) //抄送
	// m.SetHeader("Cc",
	//	m.FormatAddress("jijm@bettem.com", "收件人")) //抄送

	m.SetHeader("Subject", title) // 主题

	m.SetBody("text/html", body) // 可以放html..还有其他的
	// m.SetBody(body) // 正文
	m.Attach(filename) //添加附件

	d := gomail.NewPlainDialer("smtp.qiye.163.com", 25, "liyuan@bettem.com", "1pLvzZ9HNvcmWSRn") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return
	}

	// os.Remove(filename)
	log.Println("发送成功! 并删除excel文件成功")
}
