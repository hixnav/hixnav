# HixNav
 
#### Introduce 
Hixnav involves many languages and technologies: vue, nodejs, golang, mysql, oss, docker, json, etc. It provides a series of integrated platforms for convenient work such as custom cloud chains, cloud storage, cloud maps, and cloud accounts. Free and progressively stronger.

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
