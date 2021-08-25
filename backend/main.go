package main

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	
	"github.com/labstack/echo/v4"
)
const region = "ap-northeast-1"
const bucket = "photo-tabi-dev"

func GetFile(filename string) ([]byte, error){
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	svc := s3.New(sess)
	

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: 	aws.String(filename),
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	rc := obj.Body
	content, err := ioutil.ReadAll(rc)
	defer rc.Close()
	if err != nil { log.Fatal(err) }

	return content, nil
}

type imgGetRequest struct {
	ImgPath string `json:"img_path"`
}



func HandleImgGet(c echo.Context) error {
	var requestBody imgGetRequest
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err,
		)
	}

	if requestBody.ImgPath == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Errorf("img name is empty"),
		)
	}


	// filename := "img1.jpg"
	content, _ := GetFile(requestBody.ImgPath)
	
	enc := base64.StdEncoding.EncodeToString(content)
	
	return c.String(http.StatusOK, enc)
}


type testRequest struct {
	Sample string `json:"sample"`
}
func HandleJsonCheck(c echo.Context) error {
	var requestBody testRequest
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err,
		)
	}
	
	return c.String(http.StatusOK, requestBody.Sample)
}

func main(){
	
	fmt.Println("hello")
	
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/jsoncheck", HandleJsonCheck)
	e.GET("/trip/get_img", HandleImgGet)
	e.Logger.Fatal(e.Start(":1323"))	
}

