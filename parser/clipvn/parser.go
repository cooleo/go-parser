package clipvn

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/herotl2005/go-parser/models"
	elastigo "github.com/mattbaird/elastigo/lib"
	"golang.org/x/net/html"
	"log"
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

func GetXmlContent(Source string, Categories []string, Thumb string, Title string) []*models.VideoModel {

	var videoList []*models.VideoModel

	var model *models.VideoModel

	model = new(models.VideoModel)
	model.Image = Thumb
	model.Location = Source
	model.Title = Title
	model.Categories = Categories
	videoList = append(videoList, model)

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
	detail := document.Find("div[id=watch-region]")
	if detail != nil && detail.Length() > 0 {
		//fmt.Println("detail>>>>>>>>>>>%s",detail.First().Html())
		return true
	}
	return false
}

func GetCategory(document *goquery.Document) []string {

	var categories []string
	descriptionBox, _ := document.Find("ul.navbar-nav li.active").Html()
	if descriptionBox == "" {
		categories = append(categories, "Giải trí")
		return categories
	}
	document.Find("ul.navbar-nav li.active").Each(func(i int, s *goquery.Selection) {
		fmt.Println("category:%s", strings.TrimSpace(s.Text()))
		categories = append(categories, strings.TrimSpace(s.Text()))
	})
	return categories
}

func GetTitle(document *goquery.Document) string {
	var title = ""
	titleElement := document.Find("div.watch-region-info h1")
	if titleElement.Length() > 0 {
		title = titleElement.Text()
	}
	title = strings.TrimSpace(title)
	title = strings.Replace(title, "\t", "", -1)
	return title
}

func GetSourceList(document *goquery.Document) string {
	contentString, _ := document.Find("div[id=watch-region]").Html()
	if strings.Contains(contentString, "script") {
		startIndex := strings.LastIndex(contentString, "'file':'") + 8
		endIndex := strings.LastIndex(contentString, "','label':")
		keyString := contentString[startIndex:endIndex]
		return keyString
	}
	return ""

}
func StartParser(dao *models.Dao, document *goquery.Document, c *elastigo.Conn) bool {

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
