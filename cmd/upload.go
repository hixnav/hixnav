package cmd

import (
    "context"
    "net/http"
    "net/url"
    "fmt"

	"log"
	"mime/multipart"
	"github.com/gin-gonic/gin"
    "github.com/tencentyun/cos-go-sdk-v5"
)

var bcketMap = map[string]string{
	"1":"hixnavfile",
	"2":"hixnavnote",
	"3":"hixnavmind",
	"4":"hixnavvideo",
	"5":"hixnavmusic",
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	tencentCOS(c, "hixnav", file)
}

func UploadIO(c *gin.Context) {
	bucketId := c.PostForm("active")
	file, _ := c.FormFile("file")
	tencentCOS(c, bcketMap[bucketId], file)
}

func tencentCOS(c *gin.Context,bucket string,  file *multipart.FileHeader) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(fmt.Sprintf(COS, bucket))
	b := &cos.BaseURL{BucketURL: u} 
	cli := cos.NewClient(b, &http.Client{
	    Transport: &cos.AuthorizationTransport{
	        SecretID:  COSSecretID,   // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
	        SecretKey: COSSecretKey,  // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
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
	c.JSON(http.StatusOK, map[string]interface{}{
		"file": fmt.Sprintf(COS, "haixin")+file.Filename,
	})
}

type Bucket struct {
	Active     string `form:"active" json:"active"  binding:"required"`
}

func GetFileIO(c *gin.Context) {
	var req  Bucket
	_ = c.ShouldBindJSON(&req)
	log.Println("bucketId:",req.Active)
	tencentCOSListObj(c, bcketMap[req.Active])
}


func tencentCOSListObj(c *gin.Context,bucket string) {
    // 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
    // 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
    // COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
    u, _ := url.Parse(fmt.Sprintf(COS, bucket))
    b := &cos.BaseURL{BucketURL: u}
    cli := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
			SecretID:  COSSecretID,   // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
	        SecretKey: COSSecretKey,  // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
        },
    })


    opt := &cos.BucketGetOptions{
        Prefix: "",
        MaxKeys: 10000,
    }
    v, _, err := cli.Bucket.Get(context.Background(), opt)
    if err != nil {
		log.Println(err)
    }

	c.JSON(http.StatusOK, map[string]interface{}{
		"files": v.Contents,
	})
}