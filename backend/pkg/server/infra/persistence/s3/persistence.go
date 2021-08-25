package s3

import (
	"bytes"
	"io/ioutil"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/file"
)

const region = "ap-northeast-1"
// const bucket = "photo-tabi"
const bucket = "photo-tabi-dev"

type s3Persistence struct {
}

func NewPersistence() file.FileRepo {
	return &s3Persistence{}
}

// s3にファイルを保存する処理
func (sp s3Persistence) SaveFile(filename string, file []byte) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	// TODO: []byteとio.Readerの関係を理解する
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	})
	if err != nil {
		return "", err
	}
	return path.Base(output.Location), nil
}

// s3からファイルを取得する処理
func GetFile(filename string) ([]byte, error){
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: 	aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	
	rc := obj.Body
	content, err := ioutil.ReadAll(rc)
	defer rc.Close()
	if err != nil {
		return nil, err
	}

	return content, nil
}


func (sp s3Persistence) DeleteFile(filename string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return err
	}
	svc := s3.New(sess)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	}

	if _, err := svc.DeleteObject(input); err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return aerr
			}
		} else {
			return err
		}
	}
	return nil
}
