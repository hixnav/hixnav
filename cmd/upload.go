package cmd

import (
	"context"
	"github.com/hixnav/hixnav.git/internal/e"
	"net/http"
	"net/url"
	"strings"

	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Upload struct{}

// 目录映射：active 值对应存储桶内的目录名
var dirMap = map[string]string{
	"0": "image",
	"1": "file",
	"2": "note",
	"3": "mind",
	"4": "video",
	"5": "music",
}

// getCOSClient 创建单桶 COS 客户端
func getCOSClient() *cos.Client {
	cosURL := GlobalCOS
	if !strings.HasSuffix(cosURL, "/") {
		cosURL += "/"
	}
	u, _ := url.Parse(cosURL)
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  GlobalCOSSecretID,
			SecretKey: GlobalCOSSecretKey,
		},
	})
}

// getDir 根据 active 值获取目录前缀
func getDir(active string) string {
	if dir, ok := dirMap[active]; ok {
		return dir + "/"
	}
	return ""
}

func (s *Upload) UploadFile(c *gin.Context) (interface{}, error) {
	file, _ := c.FormFile("file")
	return uploadToCOS(c, "image", file)
}

func (s *Upload) UploadIO(c *gin.Context) (interface{}, error) {
	bucketId := c.PostForm("active")
	file, _ := c.FormFile("file")
	return uploadToCOS(c, bucketId, file)
}

func uploadToCOS(c *gin.Context, active string, file *multipart.FileHeader) (interface{}, error) {
	cli := getCOSClient()
	dir := getDir(active)
	name := dir + file.Filename

	fd, err := file.Open()
	if err != nil {
		return nil, e.AppError{Code: -1, Msg: err.Error()}
	}
	defer fd.Close()

	_, err = cli.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		return nil, e.AppError{Code: -1, Msg: err.Error()}
	}

	return map[string]interface{}{
		"file": GlobalCOS + name,
	}, nil
}

type Bucket struct {
	Active string `form:"active" json:"active"  binding:"required"`
	Key    string `form:"key" json:"key"`
}

func (s *Upload) GetFileIO(c *gin.Context) (interface{}, error) {
	var req Bucket
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return "", e.AppError{Code: -1, Msg: err.Error()}
	}
	log.Println("active:", req.Active)
	return listCOSObjects(getDir(req.Active))
}

type keyMap map[string]interface{}

func listCOSObjects(prefix string) ([]keyMap, error) {
	cli := getCOSClient()

	opt := &cos.BucketGetOptions{
		Prefix:  prefix,
		MaxKeys: 10000,
	}
	var keyMaps []keyMap
	v, _, err := cli.Bucket.Get(context.Background(), opt)
	if err != nil {
		return keyMaps, e.AppError{Code: -1, Msg: err.Error()}
	}
	for _, item := range v.Contents {
		keyMaps = append(keyMaps, keyMap{
			"key":  item.Key,
			"size": item.Size,
		})
	}
	return keyMaps, nil
}

func (s *Upload) DownFileIO(c *gin.Context) (interface{}, error) {
	var req Bucket
	_ = c.ShouldBindJSON(&req)
	log.Println("active:", req.Active)
	return downloadCOSObject(req.Key)
}

func downloadCOSObject(key string) (interface{}, error) {
	cli := getCOSClient()

	opt := &cos.ObjectGetOptions{}
	resp, err := cli.Object.Get(
		context.Background(), key, opt,
	)
	if err != nil {
		return nil, e.AppError{Code: -1, Msg: err.Error()}
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, e.AppError{Code: -1, Msg: err.Error()}
	}
	contentType := resp.Header.Get("Content-Type")
	contentLength := resp.Header.Get("Content-Length")
	resp.Body.Close()
	return map[string]interface{}{
		"file":          data,
		"contentType":   contentType,
		"contentLength": contentLength,
	}, nil
}

func (s *Upload) DelFileIO(c *gin.Context) (interface{}, error) {
	var req Bucket
	_ = c.ShouldBindJSON(&req)
	log.Println("active:", req.Active)
	return deleteCOSObject(req.Key)
}

func deleteCOSObject(key string) (interface{}, error) {
	cli := getCOSClient()

	_, err := cli.Object.Delete(context.Background(), key)
	if err != nil {
		return nil, e.AppError{Code: -1, Msg: err.Error()}
	}
	return "", nil
}
