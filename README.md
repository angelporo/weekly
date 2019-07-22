# 个人周报命令行工具, 命令行提交周报

为了避免繁琐周报, 所以实现命令行添加, `周报内容`, 和 `开始`, `结束`时间

## 使用

项目内部已编译好`window` `linux` 可执行文件, `clone`下来后请求改配置文件,

### example

`weekly 2019-07-02 2019-07-15 这里是周报纪要`

[163邮箱开启授权码](https://help.mail.163.com/faqDetail.do?code=d7a5dc8471cd0c0e8b4b8f4f8e49998b374173cfe9171305fa1ce630d7f67ac2cda80145a1742516)

配置文件 `email.config.json` 配置文件需要和可执行命令在同目录下

```json
{
  "UserName": "liyuan@bettem.com",
  "PassWord": "密码(授权码)",
  "SendTo": "liyuan@bettem.com",
  "Host": "smtp.qiye.163.com",
  "Point": 25, <!-- 我们公司163邮箱默认 25 端口 -->
  "Auth": "李渊", <!-- 这里发送人员名称 -->
  "CopyTo": [""] <!-- 抄送人邮箱地址 -->
}
```


### 如果想要自己编译, Golang 支持交叉编译 , 请配置号环境自己编译

可以源码运行 , 也可以编译后运行
[地址](https://blog.csdn.net/panshiqu/article/details/53788067)

Mac 下编译 Linux 和 Windows 64位可执行程序

- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
- CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
