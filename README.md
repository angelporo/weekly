# 个人周报命令行工具

为了避免繁琐周报, 所以实现命令行添加周报内容和开始结束时间

使用需要 163 邮箱开启 pop/smtp 密码使用授权码


# 配置文件 `email.config.json` 配置文件需要和可执行命令在同目录下

`window`下没有做测试


```
{
  "UserName": "liyuan@bettem.com",
  "PassWord": "密码(授权码)",
  "SendTo": "liyuan@bettem.com",
  "Host": "smtp.qiye.163.com",
  "Point": 25,
  "Auth": "李渊", <!--这里发送人员名称 -->
  "CopyTo": [""] <!-- 抄送人邮箱地址 -->
}
```


### Golang 支持交叉编译

可以源码运行 , 也可以编译后运行
[地址](https://blog.csdn.net/panshiqu/article/details/53788067)

Mac 下编译 Linux 和 Windows 64位可执行程序

- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
- CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
