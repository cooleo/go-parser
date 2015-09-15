package models

import (
	"gopkg.in/mgo.v2"
)

const (
	DbName          = "clipdb"
	PageCollection  = "pages"
	LinkCollection  = "links"
	TopicCollection = "topics"
	MediaCollection = "medias"
	FeedCollection  = "feeds"
    VideoCollection = "videos"
    ImageCollection = "images"

	BaseYear = 2014
)

type Dao struct {
	session *mgo.Session
}

func NewDao(Host string) (*Dao, error) {
	session, err := mgo.Dial(Host)
	if err != nil {
		return nil, err
	}
	return &Dao{session}, nil
}

func (d *Dao) Close() {
	d.session.Close()
}
