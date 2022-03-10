package model

type Admin struct {
}

func (a Admin) UserInfo(username, password string) int64 {
	var uid int64
	if username == "wenmu" {
		uid = 10000
	} else {
		uid = 0
	}
	return uid
}