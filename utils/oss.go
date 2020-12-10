package utils

import (
	"log"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"oconv/config"
)

func UploadToOss(files []string) (objs []string, err error) {
	// 创建OSSClient实例
	client, err := oss.New(config.Get().Oss.Endpoint, config.Get().Oss.AccessKey, config.Get().Oss.AccessKeySecret)
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(config.Get().Oss.Bucket)
	if err != nil {
		log.Println("[ERROR] ", err.Error())
		return
	}

	// 上传本地文件。
	for _, v := range files {
		obj := StringBuilder("res/", path.Base(v))
		if err = bucket.PutObjectFromFile(obj, v); err != nil {
			log.Println("[ERROR] ", err.Error())
			return
		}
		log.Println("[INFO] file uploaded: " + v)
		objs = append(objs, obj)
	}

	return
}
