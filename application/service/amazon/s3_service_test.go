package amazon

import (
	"log"
	"net/http"
	"testing"
)

func TestUploadFile(t *testing.T) {
	service := NewS3Service()
	opt := Opt{
		AccessKey: "",
		SecretKey: "",
		Bucket:    "pure-be-source",
		Region:    "us-west-2",
	}
	service.Init(opt)
	imageUrl := "https://s.fs128.com/data/d1dbe49d27f4744f9af852056922cdee.png"
	response, err := http.Get(imageUrl)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	key := "dev/badge/testByCharles.png"
	err = service.UploadFile(key, response.Body)
	if err != nil {
		log.Panic(err)
	}
	isResourceExist := false
	s3Object, err := service.GetObject(key)
	if err != nil {
		isResourceExist = false
	} else {
		isResourceExist = true
	}
	log.Println(isResourceExist)
	log.Println(s3Object)
}
