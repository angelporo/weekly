Golang 支持交叉编译，在一个平台上生成另一个平台的可执行程序，最近使用了一下，非常好用，这里备忘一下。
Mac 下编译 Linux 和 Windows 64位可执行程序

- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
- CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
