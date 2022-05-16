package cmd

type Account struct {
	ID       int64
	Sitename string
	Siteurl  string
	Name     string
	Password string
	Createat int64
	Uid      int64 `json:"Cate, int"`
}

func (Account) TableName() string {
	return "accounts"
}
