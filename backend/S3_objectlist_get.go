package main

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	//awsのec2に接続して、ec2内で実行
	//bucketの中身の一覧を取得, ./images/test2.jpgをアップロード、ダウンロードを行う。
	//localだと動かない。
	//以下で実行
	// $ sudo docker run -it -v $(pwd):/work -w /work golang
	// # go get -u github.com/aws/aws-sdk-go
	// # go build S3_objectlist_get
	// # ./S3_objectlist_get
	//
	//以下のurlにdownload方法など書いてある。
	//https://qiita.com/nightswinger/items/df6050ea8f8f541360f4

	sess, _ := session.NewSession(&aws.Config{//セッション確立
		Region: aws.String("ap-northeast-1")},
	)
	svc := s3.New(sess)

	bucket := "photo-tabi"//今回使うバケット名

	
	//S3の中のオブジェクト取得
	resp, _ := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})

	fmt.Println("Buckets:")
	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}	



	//ファイルのアップロード
	filename := "./images/test2.jpg"
	file, err := os.Open(filename)

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key: aws.String(filename),
		Body: file,
	})
	fmt.Println(err)
	file.Close()



	//ファイルのダウンロード
	download_file, err := os.Create("./images/downloaded_image.jpg")
	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(download_file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})
	fmt.Printf("DownloadedSize: %d byte", numBytes)
	file.Close()
}
