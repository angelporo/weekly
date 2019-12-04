# <一键发送周报> 为了大家方便使用,已经编译好`mac`下和`win`的可执行二进制文件

# 使用

根据以下数据结构 填写对应信息
```golang
type Mail struct {
	FileNameDir       string   // 文件绝对路径
	Title             string   // 标题
	UserName          string   // 用户名
	Point             int      // 开放端口
	Auth              string   // 作者 -> 邮件别名 对应excel中作者
	PassWord          string   // 密码 -> 这里需要 smtp 授权码
	SendTo            string   // 发送目标
	CopyTo            []string // 抄送目标
	Host              string   // 发送目标域
	StartTime         string   // 周报开始时间 栗子: 2019-05-12
	EndTime           string   // 周报结束时间 栗子: 2019-05-12
	Content           string   // 周报主要内容
	NextWeeklyContent string   // 下周内容
}
```

[163邮箱开启授权码(以上PassWord获取方法)](https://help.mail.163.com/faqDetail.do?code=d7a5dc8471cd0c0e8b4b8f4f8e49998b374173cfe9171305fa1ce630d7f67ac2cda80145a1742516)

- 修改`email.config.json`中对应内容
- 执行`./weekly`

# 个人周报命令行工具, 命令行提交周报


### 如果想要自己编译, Golang 支持交叉编译 , 请配置号环境自己编译

可以源码运行 , 也可以编译后运行
[地址](https://blog.csdn.net/panshiqu/article/details/53788067)

Mac 下编译 Linux 和 Windows 64位可执行程序

- CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
- CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
