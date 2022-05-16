package cmd

const (
	defaultUsername   = "hixnav"
	defaultUserPasswd = "hixnav"
)

type User struct {
	Id       int64
	Account  string
	Passwd   string
	CreateAt int64 `json:"create_at, int"`
}

func (User) TableName() string {
	return "users"
}
