![logo](http://hixnav.github.io/img/logo.png)

Hixnav involves many languages and technologies: vue, nodejs, golang, mysql, oss, docker, json, etc. It provides a series of integrated platforms for convenient work such as custom cloud chains, cloud storage, cloud maps, and cloud accounts. Free and progressively stronger.

海芯导航涉及多种语言和技术：vue、nodejs、golang、mysql、oss、docker、json等。它提供了一系列集成平台，可以方便地进行定制云链、云存储、云地图和云账户等工作。自由且逐渐强大。



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
