# HixNav
 
#### Introduce 
Hixnav supports custom navigation, custom links, cloud disk, cloud image, cloud account management

#### Config System Env

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
