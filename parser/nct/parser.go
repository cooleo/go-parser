package nct

import (
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/herotl2005/go-parser/models"
	elastigo "github.com/mattbaird/elastigo/lib"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Query struct {
	Tracklist []Video `xml:"track>item"`
}

type Video struct {
	Title    string `xml:"title"`
	Image    string `xml:"image"`
	Location string `xml:"location"`
}

func GetXmlContent(Url string, categories []string, Title string, Thumb string) []*models.VideoModel {
	resp, err := http.Get(Url)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var q Query
	xml.Unmarshal(body, &q)

	var videoList []*models.VideoModel
	for _, video := range q.Tracklist {
		if video.Location != "" {
			var model *models.VideoModel

			model = new(models.VideoModel)
			model.Image = video.Image
			model.Location = video.Location
			model.Title = strings.TrimSpace(video.Title)
			model.Categories = categories
			videoList = append(videoList, model)
		}
	}
	return videoList
}

func GetThumb(doc *goquery.Document) string {
	var thumb = ""
	metas := doc.Find("meta")
	if metas.Length() > 0 {

		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			op, _ := s.Attr("property")
			con, _ := s.Attr("content")
			if op == "og:image" {
				thumb = con
			}
		})
	}
	return thumb
}
func IsDetailPage(document *goquery.Document) bool {
	detail := document.Find("div[id=player]")
	if detail != nil && detail.Length() > 0 {
		return true
	}
	return false
}

func GetCategory(document *goquery.Document) []string {

	var categories []string
	descriptionBox, _ := document.Find("div[class='info-detail player-detail-width']").Html()
	if descriptionBox == "" {
		categories = append(categories, "Giải trí")
		return categories
	}
	document.Find("div[class='info-detail player-detail-width'] a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.Contains(href, "http://v.nhaccuatui.com/") {
			fmt.Println("category:%s", s.Text())

			categories = append(categories, strings.TrimSpace(s.Text()))
		}
	})
	return categories
}

func GetSourceList(document *goquery.Document) string {

	key, _ := document.Find("a[play_key]").First().Attr("play_key")
	fmt.Println("detail.Length():%s", key)
	if key != "" {
		return "http://v.nhaccuatui.com/flash/xml?key=" + key
	} else {
		contentString, _ := document.Html()
		if strings.Contains(contentString, "key2=") {
			startIndex := strings.Index(contentString, "key2=") + 5
			endIndex := startIndex + 32
			keyString := contentString[startIndex:endIndex]
			return "http://v.nhaccuatui.com/flash/xml?key2=" + keyString
		}
		return ""
	}
}

func StartParser(dao *models.Dao, document *goquery.Document, c *elastigo.Conn) bool {

	result := IsDetailPage(document)
	if result {

		categories := GetCategory(document)
		fmt.Println("categories :%s", categories)
		sourceUrl := GetSourceList(document)
		fmt.Println("srouceUrl:%s", sourceUrl)
		var videoList []*models.VideoModel
		videoList = GetXmlContent(sourceUrl, categories, "", "")

		topics := dao.CreateTopicFromList(categories)

		fmt.Println("topics:%s", topics)

		for _, video := range videoList {
			fmt.Println("title:%s", video.Title)
			fmt.Println("image:%s", video.Image)
			fmt.Println("location:%s", video.Location)
			fmt.Println("categories:%s", video.Categories)

			imageBucket, imageFile := models.UploadImage(video.Image)
			fmt.Println("bucket:%s, fileName:%s", imageBucket, imageFile)

			image, _ := dao.CreateImageWithParams(imageBucket, imageFile, video.Title)
			fmt.Println("image:", image)

			bucket, fileName, duration, timescale, durstr := models.UploadVideo(video.Location)
			fmt.Println("bucket:%s, fileName:%s, duaration:%d, timescale:%d, duration:%s", bucket, fileName, duration, timescale, durstr)

			video, _ := dao.CreateVideoWithParams(fileName, bucket, duration, timescale, durstr, video.Title)
			fmt.Println("video:", video)

			var videos []models.Video
			videos = append(videos, *video)
			feed, _ := dao.CreateWithParams(video.Title, "", "", imageFile, imageBucket, categories, topics, videos, *image, c)
			fmt.Println("feed:", feed)
		}
		return true
	} else {
		return false
	}

}
func Print(dao *models.Dao, page models.PageHtml, c *elastigo.Conn) {

	doc, err := html.Parse(strings.NewReader(page.Html))
	if err != nil {
		log.Fatal(err)
	}
	document := goquery.NewDocumentFromNode(doc)
	if document == nil {
	}

	StartParser(dao, document, c)

	fmt.Println("url:%s", page.Url)
}
