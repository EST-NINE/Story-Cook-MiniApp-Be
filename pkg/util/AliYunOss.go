package util

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/ncuhome/story-cook/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func AliOss(fileName string, file *multipart.FileHeader) (string, error) {
	// 创建OSSClient实例
	client, err := oss.New(config.AliEndPoint, config.AliAccessKeyId, config.AliAccessKeySecret)
	if err != nil {
		return "", err
	}

	// 获取存储空间
	bucket, err := client.Bucket(config.AliBucketName)
	if err != nil {
		return "", err
	}

	// 得到文件数据
	fileData, err := file.Open()
	defer func(fileData multipart.File) {
		err := fileData.Close()
		if err != nil {
			log.Println(err)
		}
	}(fileData)

	// 上传文件
	err = bucket.PutObject(fileName, fileData)
	if err != nil {
		return "", err
	}

	// 获取文件访问地址
	imagePath := "https://" + config.AliBucketName + "." + config.AliEndPoint + "/" + fileName
	fmt.Println("文件上传到：", imagePath)
	return imagePath, nil
}
