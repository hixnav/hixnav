# HixNav
 
#### Introduce 
You can use it as a convenient tool, or you can deploy it to the cloud and access your services anytime, anywhere, so start trying

#### Config Env

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

#### Build

```shell
# windows
$ go build -o hixnav.exe main.go

# linux
$ go build -o hixnav main.go
```

#### Run

```shell
$ hixnav
```