package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	//awsのec2に接続して、ec2内で実行
	//bucketの中身の一覧を取得
	//localだとエラーが起こりそう。
	//以下のurlにdownload方法など書いてある。
	//https://qiita.com/nightswinger/items/df6050ea8f8f541360f4

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	svc := s3.New(sess)

	bucket := "photo-tabi"//今回使うバケット名
	resp, _ := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})

	fmt.Println("Buckets:")
	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}	
}
