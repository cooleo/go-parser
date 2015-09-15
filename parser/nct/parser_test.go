package nct

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"testing"
)

func TestGetXmlContent(t *testing.T) {
	var videoList []*models.VideoModel
	videoList = GetXmlContent("http://v.nhaccuatui.com/flash/xml?key5=a1301b5cf9dffc64c8b99e5f6ee21437", nil, "", "")
	if videoList == nil {
		t.Error("Expected 1.5, got ", videoList)
	}
}

func TestIsDetailPageVideoSuccess(t *testing.T) {
	url2 := "http://v.nhaccuatui.com/video/oan.htbHSPuRJfB5n.html"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if !result {
		t.Error("Expected 1.5, got ", result)
	}
	fmt.Println("result:", result)
}

func TestIsDetailPageTvShowSuccess(t *testing.T) {
	url2 := "http://v.nhaccuatui.com/tv-show/shinee-shinee-world-the-1st-concert-in-japan.VWCAIghsUOv5.html?key=mjRpyaM3it32"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if !result {
		t.Error("Expected 1.5, got ", result)
	}
	fmt.Println("result:", result)
}

func TestIsDetailPageFail(t *testing.T) {
	url2 := "http://v.nhaccuatui.com/clip-giai-tri-video-hot-nhat.html"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if !result {
		t.Error("Expected 1.5, got ", result)
	}
	fmt.Println("result:", result)
}

func TestIsDetailPage(t *testing.T) {
	url2 := "http://v.nhaccuatui.com/tv-show/shinee-shinee-world-the-1st-concert-in-japan.VWCAIghsUOv5.html?key=mjRpyaM3it32"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if result {
		hasCategory := GetCategory(doc)
		fmt.Println("hasCategory :%s", hasCategory)
		sourceUrl := GetSourceList(doc)
		fmt.Println("srouceUrl:%s", sourceUrl)
		// t.Error("Expected 1.5, got ", result)
	}
	fmt.Println("result:", result)
}

func TestStartParser(t *testing.T) {
	url2 := "http://v.nhaccuatui.com/show-am-nhac/ca-de-dai.yvqfLo3gSzQZ.html?key=2661gCJIwT2YR"
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
	url2 := "http://v.nhaccuatui.com/tv-show/shinee-shinee-world-the-1st-concert-in-japan.VWCAIghsUOv5.html?key=mjRpyaM3it32"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if result {
		categories := GetCategory(doc)
		fmt.Println("categories :%s", categories)
		sourceUrl := GetSourceList(doc)
		fmt.Println("srouceUrl:%s", sourceUrl)
		var videoList []*models.VideoModel
		videoList = GetXmlContent(sourceUrl, categories, "", "")
		if videoList == nil {
			t.Error("Expected 1.5, got ", videoList)
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
