# go-githubhost
## 国内github访问过慢问题，项目使用Go实现，支持获取Github相关域名最快的IP，提供dns刷新相关信息

# 下载项目
```go
git clone https://github.com/w9624/go-githubhost
```
##二进制文件
**[Linux](https://imgs-1256077501.cos.ap-beijing.myqcloud.com/githubhost_linux)**
**[MacOS](https://imgs-1256077501.cos.ap-beijing.myqcloud.com/githubhost_drawin)**
**[Windows](https://imgs-1256077501.cos.ap-beijing.myqcloud.com/githubhost_windows.exe)**


# 项目运行
## 直接运行output目录中二进制文件
```shell
chmod a+x ./output/githubhost_xxx
./output/githubhost_xxx
```

## 项目源文件编译后运行
```shell
CGO_ENABLED=0 GOOS=darwin/linux/windows GOARCH=amd64 go build main.go -o ./output/githubhost_xxx
# ... 
```

## 使用Go命令运行
```shell
go run main.go
```

# 项目使用
### 浏览器访问访问：http://localhost:9090
![示例图片](https://imgs-1256077501.cos.ap-beijing.myqcloud.com/go-githubhost.png)

# 如果有用，亲，记得给个star🌟哈～