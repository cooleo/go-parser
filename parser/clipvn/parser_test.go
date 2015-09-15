package clipvn

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/herotl2005/go-parser/models"
	"testing"
)

func TestIsDetailPageVideoSuccess(t *testing.T) {
	url2 := "http://clip.vn/watch/Nhung-va-cham-nay-lua-trong-cac-cuoc-doi-dau-M-U-Liverpool,RK8-/"
	doc, _ := goquery.NewDocument(url2)
	result := IsDetailPage(doc)
	if !result {
		t.Error("Expected 1.5, got ", result)
	}
	fmt.Println("result:", result)
}

func TestIsDetailPage(t *testing.T) {
	url2 := "http://clip.vn/watch/Nhung-va-cham-nay-lua-trong-cac-cuoc-doi-dau-M-U-Liverpool,RK8-/"
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
	url2 := "http://clip.vn/watch/25-game-de-choi-tren-iPhone-6s,RKoV/"
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
	url2 := "http://clip.vn/watch/25-game-de-choi-tren-iPhone-6s,RKoV/"
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
