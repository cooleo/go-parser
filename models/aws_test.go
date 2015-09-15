package models

import (
	"fmt"
	"testing"
)

/*
func TestUploadImage(t *testing.T) {
	bucket, filename := UploadImage("http://avatar.nct.nixcdn.com/slideshow/2015/09/05/2/4/c/2/1441426162993.jpg")
	if bucket == "" && filename == "" {
		t.Error("Expected uploaded ", bucket)
	} else {
		t.Log("TestUploadImage passed.")
	}
	fmt.Println("bucket:%s, filename:%s", bucket, filename)
}

func TestUploadVideo(t *testing.T) {
	bucket, fileName, duration, timescale,durstr := UploadVideo("http://vredir.nixcdn.com/3d900b96b8ac86b0d0534f4efa861a8a/55f53674/SongClip11/ShineeWorldShineeThe1StConcertIn_46vc8.mp4")
	if bucket == "" && fileName == "" {
		t.Error("Expected uploaded ", bucket)
	} else {
		t.Log("TestUploadVideo passed.")
	}
    fmt.Println("bucket: %s, filename:%s, duration:%d, timescale:%d, duration:%s", bucket, fileName, duration, timescale,durstr)
}*/

func TestGetDuration(t *testing.T) {
	duration := GetDuration(17678000 / 1000)
	if duration == "" {
		t.Error("TestGetDuration fail ", duration)
	} else if duration == "04:54:38" {
		t.Log("TestGetDuration passed.")
	}
	fmt.Println("duration:", duration)
}
func TestExtractMp4Info(t *testing.T) {
   // dir := "/Users/hungnguyendang/work/src/go-parser-engine/models/1iWznsARxTE.mp4"
    //dir := "/Users/hungnguyendang/work/src/go-parser-engine/parser/clipvn/1iWxyKHl7PJ.mp4"
    dir := "/Users/hungnguyendang/Desktop/6c07810750a06cda6fa20db03d2d5449.mp4"
    duration, timescale, durationstr, width, height := ExtractMp4Meta(dir)
	
    fmt.Println("duaration:%d, timescale:%d, durationstr:%s, width:%d, height:%d",duration, timescale, durationstr, width, height)
  
}
