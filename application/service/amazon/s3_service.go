package amazon

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type S3Service struct {
	opt    *Opt
	client *s3.Client
}

type Opt struct {
	Region    string // 地区
	Bucket    string // 文件库
	AccessKey string // 账号
	SecretKey string // 密码
}

func NewS3Service() *S3Service {
	return &S3Service{}
}

func (s *S3Service) Init(opt Opt) {
	s.opt = &opt
	// 使用 AWS 憑證
	cred := config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
		opt.AccessKey,
		opt.SecretKey,
		"",
	))
	// 載入 AWS 設定
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(opt.Region), cred)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// 創建 S3 客戶端
	client := s3.NewFromConfig(cfg)
	s.client = client

	imageUrl := "https://s.fs128.com/data/d1dbe49d27f4744f9af852056922cdee.png"
	response, err := http.Get(imageUrl)
	if err != nil {
		log.Panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
}

// UploadFile 上傳檔案
func (s *S3Service) UploadFile(key string, body io.Reader) error {
	uploader := manager.NewUploader(s.client)
	ctxType := ContextType(key)
	uploadInput := &s3.PutObjectInput{
		Bucket:      aws.String(s.opt.Bucket),
		Key:         aws.String(key),
		ContentType: aws.String(ctxType),
		Body:        body,
	}
	result, err := uploader.Upload(context.TODO(), uploadInput)
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
		return err
	}
	fmt.Printf("file uploaded to %s success.\n", result.Location)
	return nil
}

// DownFile 下载指定文件
func (s *S3Service) DownFile(key string, filenamePath string) error {
	file, err := os.Create(filenamePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	downloader := manager.NewDownloader(s.client)
	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(s.opt.Bucket),
		Key:    aws.String(key),
	})
	return err
}

// GetObject 下载指定文件
func (s *S3Service) GetObject(key string) (*s3.GetObjectOutput, error) {
	return s.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.opt.Bucket),
		Key:    aws.String(key),
	})
}

func ContextType(fileName string) string {
	fileName = strings.ToLower(fileName)
	if strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".jpeg") {
		return "image/jpeg"
	} else if strings.HasSuffix(fileName, ".png") {
		return "image/png"
	} else if strings.HasSuffix(fileName, ".html") || strings.HasSuffix(fileName, ".htm") {
		return "text/html"
	} else if strings.HasSuffix(fileName, ".css") {
		return "text/css"
	} else if strings.HasSuffix(fileName, ".js") {
		return "application/javascript"
	} else if strings.HasSuffix(fileName, ".json") {
		return "application/json" // text/plain
	} else {
		return "application/octet-stream" // 下载
	}
}
