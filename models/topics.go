package models

import (
	"github.com/herotl2005/go-parser/common"
	"gopkg.in/mgo.v2/bson"
)

type Topic struct {
	Id   bson.ObjectId
	Name string
	Slug string
}

func (dao *Dao) CreateTopic(topic *Topic) error {
	topicCollection := dao.session.DB(DbName).C(TopicCollection)
	_, err := topicCollection.Upsert(bson.M{"id": topic.Id}, topic)
	if err != nil {
	}
	return err
}
func (dao *Dao) FindTopicBySlug(Slug string) *Topic {

	topicCollection := dao.session.DB(DbName).C(TopicCollection)
	topic := new(Topic)
	query := topicCollection.Find(bson.M{"slug": Slug})
	query.One(topic)
	return topic

}
func (dao *Dao) CreateTopicFromList(Categories []string) []Topic {
	var topics []Topic
	topicCollection := dao.session.DB(DbName).C(TopicCollection)
	for index := range Categories {
		topicName := Categories[index]
		slugName := common.GenerateTextSlug(topicName)
		existedTopic := dao.FindTopicBySlug(slugName)

		if (slugName != "" && existedTopic == nil) || (existedTopic.Id == "") {
			topic := new(Topic)
			topic.Id = bson.NewObjectId()
			topic.Name = Categories[index]
			topic.Slug = common.GenerateTextSlug(topic.Name)
			_, err := topicCollection.Upsert(bson.M{"id": topic.Id}, topic)
			if err != nil {
			}
			topics = append(topics, *topic)
		} else {
			topics = append(topics, *existedTopic)
		}

	}
	return topics

}
func (dao *Dao) FindTopics() []Topic {
	topicCollection := dao.session.DB(DbName).C(TopicCollection)
	topics := []Topic{}
	query := topicCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&topics)
	return topics
}

func (dao *Dao) FindTopicById(id string) *Topic {
	topicCollection := dao.session.DB(DbName).C(TopicCollection)
	topic := new(Topic)
	query := topicCollection.Find(bson.M{"id": id})
	query.One(topic)
	return topic
}
