package models

import (
	"fmt"
	elastigo "github.com/mattbaird/elastigo/lib"
)

type MyFeed struct {
	FeedId string
	Title  string
	Thumb  string
	Bucket string
	Topics []string
}

func IndexFeed(Feed *Feed, c *elastigo.Conn) bool {
	_, err := c.Index("feedindex", "feed", GenerateObjectId(), nil, MyFeed{Feed.Id.Hex(), Feed.Title, Feed.Thumb, Feed.ThumbBucket, Feed.Topics})
	if err != nil {
	}
	return true
}

func SearchFeed(Text string, c *elastigo.Conn) string {
	searchJson := `{
	    "query" : {
	        "term" : { "Title" : "" }
	    }
	}`
	out, _ := c.Search("feedindex", "feed", nil, searchJson)

	if len(out.Hits.Hits) == 1 {
		fmt.Println("%v", out.Hits.Hits[0].Source)
		return "true"
	}
	return "false"
}
