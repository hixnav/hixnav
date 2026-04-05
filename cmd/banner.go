package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/errcode"
	"github.com/hixnav/hixnav.git/internal/storage"
)

// SaveBanner 保存 Banner 配置到数据库
func SaveBanner(c *gin.Context) (interface{}, error) {
	var req struct {
		NavBanner  string `json:"navBanner"`
		LinkBanner string `json:"linkBanner"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, e.AppError{Code: errcode.ERROR, Msg: "invalid params"}
	}

	db := storage.GetDB()
	if db == nil {
		return nil, e.AppError{Code: errcode.ERROR, Msg: "db not initialized"}
	}

	// 确保 configs 表存在
	db.AutoMigrate(&Config{})

	// 保存或更新
	configs := []struct {
		key   string
		value string
	}{
		{"nav_banner", req.NavBanner},
		{"link_banner", req.LinkBanner},
	}

	for _, cfg := range configs {
		var existing Config
		err := db.Table("configs").Where("config_key = ?", cfg.key).First(&existing).Error
		if err != nil {
			// 不存在则插入
			db.Table("configs").Create(&Config{ConfigKey: cfg.key, Value: cfg.value})
		} else {
			// 存在则更新
			db.Table("configs").Where("config_key = ?", cfg.key).Update("value", cfg.value)
		}
	}

	return map[string]interface{}{
		"success": true,
		"message": "Banner 配置已保存",
	}, nil
}

// GetBanner 获取 Banner 配置
func GetBanner(c *gin.Context) (interface{}, error) {
	db := storage.GetDB()
	if db == nil {
		return nil, e.AppError{Code: errcode.ERROR, Msg: "db not initialized"}
	}

	db.AutoMigrate(&Config{})

	var configs []Config
	db.Table("configs").Where("config_key IN ?", []string{"nav_banner", "link_banner"}).Find(&configs)

	result := map[string]string{
		"navBanner":  "",
		"linkBanner": "",
	}
	for _, cfg := range configs {
		switch cfg.ConfigKey {
		case "nav_banner":
			result["navBanner"] = cfg.Value
		case "link_banner":
			result["linkBanner"] = cfg.Value
		}
	}

	return result, nil
}
