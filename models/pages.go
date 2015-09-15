package models

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type PageHtml struct {
	Id          bson.ObjectId
	Url         string
	Html        string
	HostName    string
	Title       string
	Thumb       string
	Description string
	Type        int
	Parsed      bool
}

func (dao *Dao) CreatePage(page *PageHtml) error {
	pageCollection := dao.session.DB(DbName).C(PageCollection)
	_, err := pageCollection.Upsert(bson.M{"id": page.Id}, page)
	if err != nil {
	}
	return err
}

func (dao *Dao) FindPages() []PageHtml {
	pageCollection := dao.session.DB(DbName).C(PageCollection)
	pages := []PageHtml{}
	query := pageCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&pages)
	return pages
}

func (dao *Dao) FindParsePages() []PageHtml {
	pageCollection := dao.session.DB(DbName).C(PageCollection)
	pages := []PageHtml{}
	query := pageCollection.Find(bson.M{"parsed":false}).Limit(5)
	query.All(&pages)
	return pages
}

func (dao *Dao) FindPageById(id string) *PageHtml {
	pageCollection := dao.session.DB(DbName).C(PageCollection)
	page := new(PageHtml)
	query := pageCollection.Find(bson.M{"id": bson.ObjectIdHex(id)})
	query.One(page)

	return page
}
func (dao *Dao) UpdatePage(Page PageHtml) bool {
	fmt.Println("Page.Id:", Page.Id)
	pageCollection := dao.session.DB(DbName).C(PageCollection)
	//err := pageCollection.Update(bson.M{"id": Page.Id}, Page)
	err := pageCollection.Update(bson.M{"id": Page.Id}, Page)
	//err := pageCollection.Update(bson.M{"_id": Page.Id}, Page)
	if err != nil {
		fmt.Println("Error Page.Id:", Page.Id)
	}
	return true


}
