package mail

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

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

func (M *Mail) getHtml() string {
	text := `<div id="qm_con_body"><div id="mailContentContainer" class="qmbox qm_con_body_content qqmail_webmail_only" style=""><div style="line-height:1.7;color:#000000;font-size:14px;font-family:Arial"><br><br><br><br><br><div style="position:relative;zoom:1">--<br><div><span style=" font-size: 13px ; ;; ">` + M.Auth + ` | 软件开发部</span></div><div><span style=" font-size: 13px ; ;; "><span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="176-3669-9981">176-3669-9981</span> | <a href="mailto:` + M.UserName + `" src="mailto:guojia@bettem.com" rel="noopener" target="_blank">` + M.UserName + `</a></span></div><div><br></div><div><span style=" font-size: 13px ; ;; "><b>山西百得科技开发股份有限公司</b></span></div><div><span style=" font-size: 13px ; ;; ">股票代码 | 839289</span></div><div><span style="  ;; font-size: 13px; ">公司地址：山西 | 太原 南中环街401号数码港A座二层</span></div><div><span style="  ;; font-size: 13px; ">公司网址：<a>www.bettem.com</a> </span></div><div><span style="  ;; font-size: 13px; ">公司电话：<span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="0351-7033691">0351-7033691</span></span></div><div style="clear:both"></div></div><div id="dvLetterAngle"> </div></div><style type="text/css">.qmbox style, .qmbox script, .qmbox head, .qmbox link, .qmbox meta {display: none !important;}</style></div></div>`
	return `
    <html>
    <body>
    <h3>
    ` + text + `
    </h3>
    </body>
    </html>
    `
}

func GetConfig() (Mail, error) {
	var config Mail
	data, readErr := ioutil.ReadFile("./email.config.json")
	if readErr != nil {
		return Mail{}, errors.New("请检查你的配置文件")
	}
	unmarErr := json.Unmarshal(data, &config)
	if unmarErr != nil {
		return Mail{}, unmarErr
	}
	return config, nil
}

type Cc struct {
	Email string
	Alias string
	Type  string // Cc 抄送目标 To 发送目标
}

func (M *Mail) Send() error {
	m := gomail.NewMessage()

	var auth string
	if M.Auth == "" {
		// 默认发送作者名称
		auth = "李渊"
	} else {
		auth = M.Auth
	}
	m.SetAddressHeader("From", M.UserName /*"发件人地址"*/, auth) // 发件人

	m.SetHeader("To",
		m.FormatAddress(M.SendTo, "收件人")) // 收件人

	m.SetHeader("Cc",
		M.CopyTo...) // 抄送

	m.SetHeader("Subject", M.Title) // 主题

	m.SetBody("text/html", M.getHtml()) // 可以放html..还有其他的
	m.Attach(M.FileNameDir)             //添加附件

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
