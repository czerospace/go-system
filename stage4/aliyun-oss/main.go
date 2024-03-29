package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// Endpoint 以河源为例，其它 Region 请按实际情况填写。
	endpoint := "oss-cn-heyuan.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := ""
	accessKeySecret := ""
	bucketName := ""
	// <yourObjectName> 上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如 abc/efg/123.jpg。
	// 斜线前面代表路径
	objectName := "niko/first.jpg"
	// <yourLocalFileName> 由本地文件路径加文件名包括后缀组成，例如 /users/local/myfile.txt。
	localFileName := `C:\Users\czero\Pictures\hitgirl.webp`
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}
