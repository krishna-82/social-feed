package dao

import (
	"log"

	. "github.com/krishna-82/social-feed/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// database feed on server
type FeedDAO struct {
	Server   string
	Database string
}

// collection Follow
type Follow struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	follow_to     int
	follow_by     int
	Timestamp 	time.Time
}

// collection Feed
type Feed struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	post_by     int
	post_type   string
	post_text	string
	Timestamp 	time.Time
}

// collection User
type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	first_name 	string
	last_name   string
	address		string
	Timestamp 	time.Time
}

var db *mgo.Database

const (
	COLLECTION_FEED = "feed"
	COLLECTION_FOLLOW = "follow"
)


// Establish a connection to database
func (m *FeedDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of people user follows
func (m *follow) GetFollowTo(user_id int) ([]follow, error) {
	var results []Follow
	err := db.C(COLLECTION_FOLLOW).Find(bson.M{"follow_by": user_id}).SELECT(bson.M{"follow_to"}).All(&results)
	return results, err
}

// Find feeds for user
func (m *FeedDAO) GetFeeds([]followTo, limit int) ([]Feeds, error) {
	var feeds []Feed
	err := db.C(COLLECTION_FEED).Find(bson.M{post_by: bson.M{"$in": followTo}}}).Sort("-timestamp").Limit(limit).All(&feed)
	return feeds, err
}

// Find search result for user based on keyword
func (m *FeedDAO) Search(keyword string, limit int) ([]searchResult, error) {
	var searchResult []feeds
	
	pipeline := []bson.M{ 
		bson.M{"$match": bson.M{"$or": [ bson.M{"post_text": bson.M{keyword}},
										bson.M{"$user.first_name": bson.M{keyword}}, 
										bson.M{"$user.last_name": bson.M{keyword}}, 
										bson.M{"$user.address": bson.M{"/{keyword}/"}}
										]
									},
		bson.M{"$lookup": bson.M{ 	"from": "User",
									"localField": "post_by",
									"foreignField": "ID",
									"as": "user"},
		bson.M{"$unwind": "$user" },
		bson.M{"$sort": bson.M{ "-timestamp" }
		bson.M{"$limit": bson.M{limit}
	}}

	pipe := db.C(COLLECTION_FEED).Pipe(pipeline)
	searchResult := []bson.M{}
	err = pipe.All(&searchResult)
	return searchResult, err
}
