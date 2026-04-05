package cmd

// Config 系统配置表，用于存储 banner 等键值对配置
type Config struct {
	ID        int64
	ConfigKey string `gorm:"column:config_key"`
	Value     string
}

func (Config) TableName() string {
	return "configs"
}
