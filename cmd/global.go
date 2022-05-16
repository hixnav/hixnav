package cmd

import (
	"gorm.io/gorm"
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

	GlobalDB *gorm.DB
)
