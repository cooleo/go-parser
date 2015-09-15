package models

import (
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"time"
)

type Link struct {
	Id  string
	Url string
}

func (dao *Dao) CreateLink(link *Link) error {
	linkCollection := dao.session.DB(DbName).C(LinkCollection)	
	_, err := linkCollection.Upsert(bson.M{"id": link.Id}, link)
	if err != nil {		
	}
	return err
}
func (dao *Dao) FindLinks() []Link {
	linkCollection := dao.session.DB(DbName).C(LinkCollection)
	links := []Link{}
	query := linkCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&links)
	return links
}
func (dao *Dao) FindLinkById(id string) *Link {
	linkCollection := dao.session.DB(DbName).C(LinkCollection)
	link := new(Link)
	query := linkCollection.Find(bson.M{"id": bson.ObjectIdHex(id)})
	query.One(link)
	return link
}
