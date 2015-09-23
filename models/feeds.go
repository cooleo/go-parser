package models

import (
	elastigo "github.com/mattbaird/elastigo/lib"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Feed struct {
	Id          bson.ObjectId
	Url         string
	HostName    string
	Title       string
	Thumb       string
	ThumbBucket string
	Description string
	Type        uint32
	Topics      []string
	Photo       Image
	Videos      []Video
	TopicsList  []Topic
	Created     time.Time
	Updated     time.Time
}

func (dao *Dao) CreateFeed(feed *Feed) error {
	feedCollection := dao.session.DB(DbName).C(FeedCollection)
	_, err := feedCollection.Upsert(bson.M{"_id": feed.Id}, feed)
	if err != nil {
	}
	return err
}
func (dao *Dao) FindFeedByTitle(title string) bool {
	feedCollection := dao.session.DB(DbName).C(FeedCollection)
	pages := []Feed{}
	query := pageCollection.Find(bson.M{"Title": title}}).Limit(1)
	query.All(&pages)
	if cap(pages) >= 1 {
		return true
	} else {
		return false
	}

}
func (dao *Dao) CreateWithParams(Title string, Url string, HostName string, Thumb string, ThumbBucket string, Topics []string, TopicList []Topic, Videos []Video, Photo Image, c *elastigo.Conn) (*Feed, error) {
	var feed *Feed

	feed = new(Feed)
	feed.Id = bson.NewObjectId()
	feed.Title = Title
	feed.Url = Url
	feed.HostName = HostName
	feed.Type = 0
	feed.Thumb = Thumb
	feed.ThumbBucket = ThumbBucket
	feed.Photo = Photo
	feed.Created = time.Now()
	feed.Updated = time.Now()
	feed.Topics = Topics
	feed.TopicsList = TopicList
	feed.Videos = Videos
	feedCollection := dao.session.DB(DbName).C(FeedCollection)
	_, err := feedCollection.Upsert(bson.M{"id": feed.Id}, feed)
	if err != nil {
	}
	isIndexed := IndexFeed(feed, c)
	if isIndexed {
	}
	return feed, err

}

func (dao *Dao) FindFeeds() []Feed {
	feedCollection := dao.session.DB(DbName).C(FeedCollection)
	feeds := []Feed{}
	query := feedCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&feeds)
	return feeds
}

func (dao *Dao) FindFeedById(id string) *Feed {
	feedCollection := dao.session.DB(DbName).C(FeedCollection)
	feed := new(Feed)
	query := feedCollection.Find(bson.M{"id": bson.ObjectIdHex(id)})
	query.One(feed)
	return feed
}
