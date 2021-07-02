package s3

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/file"
)

const region = "ap-northeast-1"
const bucket = "photo-tabi"

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
	fmt.Println(output.Location)
	return output.Location, nil
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
