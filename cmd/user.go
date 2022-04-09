package cmd

import (
	"encoding/base64"
)

type Admin struct {
}

type User struct {
	Id       int64
	Account  string
	CreateAt string `json:"create_at, int"`
}

func (a Admin) UserInfo(username, password string) (uid int64) {
	var u User
	err := GlobalDB.Table("users").Where("account = ? ", username).Where("passwd = ?", base64.StdEncoding.EncodeToString([]byte(password))).First(&u)
	if err != nil {
		uid = 0
	}
	return u.Id
}
