# HixNav 海芯导航
 
#### 介绍 
海芯导航涉及多种语言和技术：vue、nodejs、golang、mysql、oss、docker、json等。它提供了一系列集成平台，可以方便地进行定制云链、云存储、云地图和云账户等工作。自由且逐渐强大。

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
