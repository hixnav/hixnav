package cmd

type Link struct {
	Id       int64
	Type     int64 `json:"Type, string"`
	Catename string
	Name     string
	Logo     string
	Url      string
	Uid      int64
}

func (Link) TableName() string {
	return "links"
}
