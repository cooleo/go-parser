package tvzing

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/herotl2005/go-parser/models"
	"testing"
)

func TestIsDetailPageVideoSuccess(t *testing.T) {
	url2 := "http://tv.zing.vn/video/Thach-Thuc-Danh-Hai-Tap-12/IWZB67AF.html"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if !result {
		t.Error("Expected 1.5, got ", result)
	}
	fmt.Println("result:", result)
}

func TestIsDetailPage(t *testing.T) {
	url2 := "http://tv.zing.vn/video/Nhan-Qua-Cuoc-Doi-Trailer/IWZB86F9.html"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if result {
		hasCategory := GetCategory(doc)
		fmt.Println("hasCategory :%s", hasCategory)
		thumb := GetThumb(doc)
		fmt.Println("thumb :%s", thumb)
		title := GetTitle(doc)
		fmt.Println("title :%s", title)

		source := GetSourceList(doc)
		fmt.Println("source :%s", source)
		t.Error("Expected 1.5, got ", result)
	}
	fmt.Println("result:", result)
}

func TestStartParser(t *testing.T) {
	url2 := "http://tv.zing.vn/video/Nhan-Qua-Cuoc-Doi-Trailer/IWZB86F9.html"
	document, _ := goquery.NewDocument(url2)
	result := StartParser(document)
	if !result {
		t.Error("TestStartParser testing got ", result)
	} else {
		t.Log("TestStartParser passed!")
	}
	fmt.Println("resutl:", result)

}

func TestParserEndtoEnd(t *testing.T) {
	url2 := "http://tv.zing.vn/video/Nhan-Qua-Cuoc-Doi-Trailer/IWZB86F9.html"
	document, _ := goquery.NewDocument(url2)
	result := IsDetailPage(document)
	if result {
		categories := GetCategory(document)
		fmt.Println("categories :%s", categories)
		sourceUrl := GetSourceList(document)
		fmt.Println("srouceUrl:%s", sourceUrl)

		thumb := GetThumb(document)
		fmt.Println("thumb :%s", thumb)
		title := GetTitle(document)

		var videoList []*models.VideoModel
		videoList = GetXmlContent(sourceUrl, categories, thumb, title)
		if videoList == nil {
		}
		for _, video := range videoList {
			fmt.Println("title:%s", video.Title)
			fmt.Println("image:%s", video.Image)
			fmt.Println("location:%s", video.Location)
			fmt.Println("categories:%s", video.Categories)
		}

	}

	fmt.Println("result:", result)
}
