# HixNav 海芯导航
 
#### 介绍 
您可以将其作为一种方便的工具使用，也可以将其部署到云上并在任何时间、任何地点访问您的服务，所以开始尝试吧

#### 配置环境变量

```shell
# mysql数据库
# DNS={user}:{password}@tcp({host}:{port})/haixin?charset=utf8mb4&parseTime=True&loc=Local
DNS

# tecentCOS腾讯云存储
# COS=https://%s-{your}.myqcloud.com/
COS 
COSSecretID
COSSecretKey
``` 

#### 编译

```shell
# windows
$ go build -o hixnav.exe main.go

# linux
$ go build -o hixnav main.go
```

#### 运行

```shell
$ hixnav
```
