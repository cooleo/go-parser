package models

import (
	"gopkg.in/mgo.v2/bson"
   
)

type Media struct {
    Id bson.ObjectId 
	Filename    string
	Label       string
	Bucket      string
	Title       string	
	Thumb       string
    ThumbBucket string
	ThumbOrigin string
	VideoOrigin string
	Video       string	
	Length      uint32
	Duration    string
    Type        uint32
}

func (dao *Dao) CreateMedia(media *Media) error {
	mediaCollection := dao.session.DB(DbName).C(MediaCollection)
	_, err := mediaCollection.Upsert(bson.M{"id": media.Id}, media)
	if err != nil {		
	}
	return err
}

func (dao *Dao) FindMedias() []Media {
	mediaCollection := dao.session.DB(DbName).C(MediaCollection)
	medias := []Media{}
	query := mediaCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&medias)
	return medias
}

func (dao *Dao) FindMediaById(id string) *Media {
	mediaCollection := dao.session.DB(DbName).C(MediaCollection)
	media := new(Media)
	query := mediaCollection.Find(bson.M{"id": id})
	query.One(media)
	return media
}
