# FF-Backup
Fanfou backup tools via golang

# <del>deprecated</del> delayed
饭否无法模拟登陆，问题暂时不想解决......以后有时间再做......

# 交叉编译


```
Windows 下编译 Mac 和 Linux 64位可执行程序

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```


Mac 下编译 Linux 和 Windows 64位可执行程序

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

Linux 下编译 Mac 和 Windows 64位可执行程序

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

# 验证码解决

半成品。根据链接下载验证码，之后手动识别。

```
r := regexp.MustCompile(`<img src="\/\/(.*?)" width`)
body := string(res)
captchAddress := r.FindStringSubmatch(body)[1]
```


