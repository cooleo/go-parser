package models

import (
	"gopkg.in/mgo.v2/bson"
   
)

type Video struct {
    Id          bson.ObjectId 
    Title       string
	FileName    string
	Label       string
	Bucket      string		
	Length      uint32
    Timescale   uint32
	Duration    string
    Type        uint32
    
}

func (dao *Dao) CreateVideo(video *Video) error {
	VideoCollection := dao.session.DB(DbName).C(VideoCollection)
	_, err := VideoCollection.Upsert(bson.M{"_id": video.Id}, video)
	if err != nil {		
	}
	return err
}

func (dao *Dao) CreateVideoWithParams(FileName string, Bucket string, Length uint32, Timescale uint32, Duration string,Title string) (*Video, error) {    
    video := new (Video)
    video.Id  =  bson.NewObjectId()
    video.Title = Title
    video.Type = 0
    video.FileName = FileName
    video.Bucket = Bucket
    video.Length = Length
    video.Timescale = Timescale
    video.Duration = Duration   
    
	VideoCollection := dao.session.DB(DbName).C(VideoCollection)
	_, err := VideoCollection.Upsert(bson.M{"id": video.Id}, video)
	if err != nil {		
	}
	return video, err
}


func (dao *Dao) FindVideos() []Video {
	VideoCollection := dao.session.DB(DbName).C(VideoCollection)
	videos := []Video{}
	query := VideoCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&videos)
	return videos
}

func (dao *Dao) FindVideoById(id string) *Video {
	VideoCollection := dao.session.DB(DbName).C(VideoCollection)
	video := new(Video)
	query := VideoCollection.Find(bson.M{"id": id})
	query.One(video)
	return video
}
