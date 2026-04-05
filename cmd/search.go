package cmd

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/storage"
)

// SearchRequest 搜索请求参数
type SearchRequest struct {
	Keyword string `json:"keyword" form:"keyword"`
}

// SearchResult 搜索结果项
type SearchResult struct {
	ID       int64  `json:"id"`
	Type     string `json:"type"` // "nav" or "link"
	Cate     int64  `json:"cate,omitempty"`
	Catename string `json:"catename"`
	Name     string `json:"name"`
	Desc     string `json:"desc,omitempty"`
	Url      string `json:"url"`
	Logo     string `json:"logo"`
}

// Search 全局搜索接口，搜索 navs 和 links 表
func (s *Nav) Search(c *gin.Context) (interface{}, error) {
	var req SearchRequest
	// 支持 JSON 和 form 两种格式
	_ = c.ShouldBindJSON(&req)
	keyword := strings.TrimSpace(req.Keyword)
	if keyword == "" {
		keyword = strings.TrimSpace(c.DefaultPostForm("keyword", c.DefaultQuery("keyword", "")))
	}
	if keyword == "" {
		return map[string]interface{}{
			"navs":  []SearchResult{},
			"links": []SearchResult{},
		}, nil
	}

	like := "%" + keyword + "%"

	// 搜索 navs 表
	var navResults []SearchResult
	storage.GetDB().Table("navs").
		Select("id, cate, catename, name, `desc`, url, logo").
		Where("name LIKE ? OR `desc` LIKE ? OR url LIKE ? OR catename LIKE ?", like, like, like, like).
		Find(&navResults)
	for i := range navResults {
		navResults[i].Type = "nav"
	}

	// 搜索 links 表
	var linkResults []SearchResult
	storage.GetDB().Table("links").
		Select("id, catename, name, url, logo").
		Where("name LIKE ? OR url LIKE ? OR catename LIKE ?", like, like, like).
		Find(&linkResults)
	for i := range linkResults {
		linkResults[i].Type = "link"
	}

	return map[string]interface{}{
		"navs":  navResults,
		"links": linkResults,
	}, nil
}
