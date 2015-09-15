package models

import (

	"fmt"
	elastigo "github.com/mattbaird/elastigo/lib"
)

// var (
// 	host *string = flag.String("host", "104.197.83.34", "Elasticsearch Host")
// )



type MyFeed struct {
	FeedId string
	Title  string
	Thumb  string
	Bucket string
	Topics []string
}

func IndexFeed(Feed *Feed, c *elastigo.Conn) bool {
	// c := elastigo.NewConn()
	// log.SetFlags(log.LstdFlags)
	// flag.Parse()
	// // Trace all requests
	// c.RequestTracer = func(method, url, body string) {
	// 	log.Printf("Requesting %s %s", method, url)
	// 	log.Printf("Request body: %s", body)
	// }
	// fmt.Println("host = ", *host)
	// c.Domain = *host
	_, err := c.Index("feedindex", "feed", GenerateObjectId(), nil, MyFeed{Feed.Id.Hex(), Feed.Title, Feed.Thumb, Feed.ThumbBucket, Feed.Topics})
	if err != nil {
	}
	return true
}

func SearchFeed(Text string, c *elastigo.Conn) string {

	// c := elastigo.NewConn()
	// log.SetFlags(log.LstdFlags)
	// flag.Parse()

	// // Trace all requests
	// c.RequestTracer = func(method, url, body string) {
	// 	log.Printf("Requesting %s %s", method, url)
	// 	log.Printf("Request body: %s", body)
	// }

	// fmt.Println("host = ", *host)
	// // Set the Elasticsearch Host to Connect to
	// c.Domain = *host

	searchJson := `{
	    "query" : {
	        "term" : { "Title" : "" }
	    }
	}`
	out, _ := c.Search("feedindex", "feed", nil, searchJson)

	if len(out.Hits.Hits) == 1 {
		fmt.Println("%v", out.Hits.Hits[0].Source)
		//return out.Hits.Hits[0].Source
		return "true"
	}
	return "false"
}
