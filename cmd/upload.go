package cmd

import (
	"context"
	"fmt"
	"github.com/hixnav/hixnav.git/internal/e"
	"net/http"
	"net/url"
	"strconv"

	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Upload struct{}

var bcketMap = map[string]string{
	"0": "hixnav",
	"1": "hixnavfile",
	"2": "hixnavnote",
	"3": "hixnavmind",
	"4": "hixnavvideo",
	"5": "hixnavmusic",
}

func bucket(c *gin.Context, bucketId string) (b string) {
	uid, _ := c.Get("uid")
	log.Println("upload uid: ", uid)
	b = bcketMap[bucketId]
	if uid != int64(1) {
		b = bcketMap[bucketId] + strconv.Itoa(int(uid.(int64)))
	}
	return b
}

func (s *Upload) UploadFile(c *gin.Context) (interface{}, error) {
	file, _ := c.FormFile("file")
	return tencentCOS(c, bucket(c, "0"), file)
}

func (s *Upload) UploadIO(c *gin.Context) (interface{}, error) {
	bucketId := c.PostForm("active")
	file, _ := c.FormFile("file")
	return tencentCOS(c, bucket(c, bucketId), file)
}

func tencentCOS(c *gin.Context, bucket string, file *multipart.FileHeader) (interface{}, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(fmt.Sprintf(GlobalCOS, bucket))
	b := &cos.BaseURL{BucketURL: u}
	cli := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  GlobalCOSSecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: GlobalCOSSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := file.Filename
	fd, err := file.Open()
	if err != nil {
		log.Println(err)
	}
	defer fd.Close()
	_, err = cli.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		log.Println(err)
	}
	return map[string]interface{}{
		"file": fmt.Sprintf(GlobalCOS, "hixnav") + file.Filename,
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
	log.Println("bucketId:", req.Active)
	return tencentCOSListObj(c, bucket(c, req.Active))
}

type keyMap map[string]interface{}

func tencentCOSListObj(c *gin.Context, bucket string) ([]keyMap, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(fmt.Sprintf(GlobalCOS, bucket))
	b := &cos.BaseURL{BucketURL: u}
	cli := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  GlobalCOSSecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: GlobalCOSSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})

	opt := &cos.BucketGetOptions{
		Prefix:  "",
		MaxKeys: 10000,
	}
	var keyMaps []keyMap
	v, _, err := cli.Bucket.Get(context.Background(), opt)
	if err != nil {
		log.Println(err)
		return keyMaps, err
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
	log.Println("bucketId:", req.Active)
	return tencentCOSDownObj(c, bucket(c, req.Active), req.Key)
}

func tencentCOSDownObj(c *gin.Context, bucket, key string) (interface{}, error) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(fmt.Sprintf(GlobalCOS, bucket))
	b := &cos.BaseURL{BucketURL: u}
	cli := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  GlobalCOSSecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: GlobalCOSSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})

	opt := &cos.ObjectGetOptions{}
	resp, err := cli.Object.Get(
		context.Background(), key, opt,
	)
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%v", *resp)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
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
	log.Println("bucketId:", req.Active)
	return tencentCOSDelObj(c, bucket(c, req.Active), req.Key)
}

func tencentCOSDelObj(c *gin.Context, bucket, key string) (interface{}, error) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(fmt.Sprintf(GlobalCOS, bucket))
	b := &cos.BaseURL{BucketURL: u}
	cli := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  GlobalCOSSecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: GlobalCOSSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})

	_, err := cli.Object.Delete(context.Background(), key)
	if err != nil {
		log.Println(err)
	}
	return "", nil
}
