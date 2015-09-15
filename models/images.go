package models

import (
	"gopkg.in/mgo.v2/bson"
   
)

type Image struct {
    Id          bson.ObjectId 
    Title       string
	FileName    string	
	Bucket      string			
    Type        uint32
    
}

func (dao *Dao) CreateImage(video *Image) error {
	imageCollection := dao.session.DB(DbName).C(ImageCollection)
	_, err := imageCollection.Upsert(bson.M{"id": video.Id}, video)
	if err != nil {		
	}
	return err
}

func (dao *Dao) CreateImageWithParams(FileName string, Bucket string, Title string) (*Image, error) {
    var image *Image
    image = new (Image)
    image.Id  =  bson.NewObjectId()
    image.Title = Title
    image.Type = 1
    image.FileName = FileName
    image.Bucket = Bucket
     
	imageCollection := dao.session.DB(DbName).C(ImageCollection)
	_, err := imageCollection.Upsert(bson.M{"id": image.Id}, image)
	if err != nil {		
	}
	return image, err
}


func (dao *Dao) FindImages() []Image {
	imageCollection := dao.session.DB(DbName).C(ImageCollection)
	images := []Image{}
	query := imageCollection.Find(bson.M{}).Sort("-cdate").Limit(50)
	query.All(&images)
	return images
}

func (dao *Dao) FindImageById(id string) *Image {
	imageCollection := dao.session.DB(DbName).C(ImageCollection)
	image := new(Image)
	query := imageCollection.Find(bson.M{"_id": id})
	query.One(image)
	return image
}
