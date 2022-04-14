![logo](http://hixnav.github.io/img/logo.png)

You can use it as a convenient tool, or you can deploy it to the cloud and access your services anytime, anywhere, so start trying

海芯导航 - 你可以把它作为一个便捷工具来使用，或者你可以把它部署到云上，随时随地访问你的服务，那么开始尝试吧

## 开始使用
### 获取HixNav

获取windows版本，点击<a href="https://github.com/hixnav/hixnav/raw/master/release/hixnav-winx64.exe">下载</a>

获取Max版本，点击<a href="https://github.com/hixnav/hixnav/raw/master/release/hixnav-mac-amd64">下载</a>

获取Linux版本，点击<a href="https://github.com/hixnav/hixnav/raw/master/release/hixnav-linux-amd64">下载</a>

查看源码，请访问仓库<a href="https://github.com/hixnav/hixnav">HixNav</a>

### 配置

>你需要一个Mysql的数据库，以及一个对象存储服务，这些服务可以从云服务厂商那里获取。

###### 配置系统环境变量

```shell
# mysql
# DNS={user}:{password}@tcp({host}:{port})/haixin?charset=utf8mb4&parseTime=True&loc=Local
DNS

# tecentCOS
# COS=https://%s-{your}.myqcloud.com/
COS 
COSSecretID
COSSecretKey
``` 

### 运行

```shell
$ hixnav
```

### 访问

> 系统默认使用80端口

http://{ip}

## 反馈

邮箱：kefu@iflet.cn
