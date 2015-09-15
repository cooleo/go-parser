package models

import (
    "testing"
    "fmt"
   // "gopkg.in/mgo.v2/bson"
)


func TestIndexFeed(t *testing.T) {
    
    
    dao, err := NewDao()
	if err != nil {
	}
	defer dao.Close()
    feed := dao.FindFeedById("55f5bbe8b42a03099b000011")

    fmt.Println("feed:", feed.Title)
    
    rs := IndexFeed(feed)
    if !rs  {
        t.Error("TestIndexFeed fail.")
    } else {
        t.Log("TestIndexFeed sucess.")
    }
    fmt.Println("result:", rs)
    
}
func TestIndexNewFeed(t *testing.T) {
        
    
//    topics := []string{"Penn", "Teller"}
//    feed := new(Feed)
//    feed.Id = bson.NewObjectId()
//    feed.Title = "Hello world"
//    feed.Thumb = "thumb"
//    feed.ThumbBucket = "bucket"
//    feed.Topics = topics
    
//      rs := IndexFeed(feed)
//    if !rs  {
//        t.Error("TestIndexFeed fail.")
//    } else {
//        t.Log("TestIndexFeed sucess.")
//    }
//    fmt.Println("result:", rs)
//    
    
}