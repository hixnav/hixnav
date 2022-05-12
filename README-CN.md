# 海芯导航（HixNav）
 
### 介绍
 
海芯导航支持自定义导航、自定义链接、云盘、云图、云账号管理

Github: https://github.com/hixnav/hixnav Gitee: https://gitee.com/hixnav/hixnav

### 快速开始

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

#### 整包编译

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
