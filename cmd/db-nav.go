package cmd

type Nav struct {
	Id       int64
	Cate     int64 `json:"Cate, int"`
	Catename string
	Name     string
	Logo     string
	Desc     string
	Url      string
	Uid      int64
}

func (Nav) TableName() string {
	return "navs"
}
