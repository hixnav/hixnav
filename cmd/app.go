package cmd

import "gitee.com/wennmu/gopkg.git/doconfig"

func InitApp() error {
	//初始化数据库
	GlobalAppSecret = doconfig.GetString("AppSecret")
	return nil
}
