package models

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"	
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
    "github.com/leelynne/mp4"
)

var (
	localPath string
	bucket    string
	prefix    string
)

const (
	ImageBucket = "vi-mages"
	VideoBucket = "viet-videos"
)

func UploadImage(Url string) (string, string) {
	bucket := ImageBucket
	defaults.DefaultConfig.Region = aws.String("us-east-1")
	fileName := GenerateObjectId() + ".jpg"
	fmt.Println("Downloading", Url, "to", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return "", ""
	}
	defer file.Close()

	response, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error while downloading", Url, "-", err)
		return "", ""
	}
	defer response.Body.Close()

	n, err := io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", Url, "-", err)
		return "", ""
	}
	fmt.Println("n:%s", n)

	uploader := s3manager.NewUploader(nil)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &fileName,
		Body:   file,
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}
	log.Println("Uploaded....%s", result)
    
    dir, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dir:%s\n", dir)
    
    err = os.Remove(dir)
    if err != nil {
    }
    
	return bucket, fileName
}

func GetDuration(Duration uint32) string {
	var result uint32
	var hours, minutes, seconds uint32
	result = Duration / 3600
    hours = result
	minutes = (Duration - (hours*3600)) / 60
	seconds = (Duration - (minutes*60)) % 60
	if Duration > 3600 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else if Duration > 60 {
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	} else {
		return fmt.Sprintf("%02d:%02d", 0, result)
	}

}

func ExtractMp4Meta(Dir string) (uint32, uint32, string, int, int) {
    m, err := mp4.Open(Dir)
	if err != nil {
	  fmt.Println(".........ExtractMp4Meta Error:%s", err)
	}
    fmt.Println("video %dx%d", m.W, m.H)
    fmt.Println("duration %d", uint32(m.Dur))
    log.Println("duration time:%s", GetDuration(uint32(m.Dur)))    
    return uint32(m.Dur), 1, GetDuration(uint32(m.Dur)), m.W,  m.H
 
}

func UploadVideo(Url string) (string, string, uint32, uint32, string) {
	bucket := VideoBucket
	defaults.DefaultConfig.Region = aws.String("us-east-1")

	fileName := GenerateObjectId() + ".mp4"
	fmt.Println("Downloading", Url, "to", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return "", "", 0, 0,""
	}
	defer file.Close()

	response, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error while downloading", Url, "-", err)
		return "", "", 0, 0,""
	}
	defer response.Body.Close()

	n, err := io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", Url, "-", err)
		return "", "", 0, 0,""
	}
	fmt.Println("n:%s", n)

	dir, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println("error file Abs:%s\n", err)
	}
	fmt.Println(".......file path:%s\n", dir)
    
    duration, timescale, durationstr, width, height := ExtractMp4Meta(dir)
	
    fmt.Println("duaration:%d, timescale:%d, durationstr:%s, width:%d, height:%d",duration, timescale, durationstr, width, height)
	fmt.Println("Start upload file to S3:%s\n", fileName)
	uploader := s3manager.NewUploader(nil)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &fileName,
		Body:   file,
	})
	if err != nil {
		fmt.Println("Failed to upload", err)
	}
	fmt.Println("Uploaded....%s", result)
    err = os.Remove(dir)
    if err != nil {
    }
   // durstr := GetDuration(duration/timescale)
	return VideoBucket, fileName, duration, timescale,durationstr
}
