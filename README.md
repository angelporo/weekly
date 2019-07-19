# 个人周报命令行工具

为了避免繁琐周报, 所以实现命令行添加周报内容和开始结束时间

使用需要 163 邮箱开启 pop/smtp 密码使用授权码


## TODO

- 收件人 抄送人 个人内容 配置化


### Golang 支持交叉编译

可以源码运行 , 也可以编译后运行
[地址](https://blog.csdn.net/panshiqu/article/details/53788067)

Mac 下编译 Linux 和 Windows 64位可执行程序

- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
- CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
