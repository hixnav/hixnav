package cmd

import (
	"os"
)

var (
	// mysql db
	GlobalMysqlDNS = ""

	// tecentCOS
	GlobalCOS          = ""
	GlobalCOSSecretID  = ""
	GlobalCOSSecretKey = ""

	// 登录验证秘钥，请务必更改
	GlobalAppSecret = "wN3bdUQbBuL2AD7PeU9VYIqO7a35VU0G"
)

func Init() {
	if os.Getenv("GlobalAppSecret") != "" {
		GlobalAppSecret = os.Getenv("AppSecret")
	}
	GlobalMysqlDNS = os.Getenv("DNS")
	GlobalCOS = os.Getenv("COS")
	GlobalCOSSecretID = os.Getenv("COSSecretID")
	GlobalCOSSecretKey = os.Getenv("COSSecretKey")
}
