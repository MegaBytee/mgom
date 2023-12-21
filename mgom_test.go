package mgom

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

//testing
// when you run test first time, delete the testing database manualy in mongodb to avoid test failure for the second time because of duplicate key index
// simple example test
// TODO all testing cases with goconvey

const (
	DB_NAME = "mgom_test"
	DB_URL  = "mongodb://localhost:27017/"
)

// Mongo instance should be global pointer variable
var (
	Mg    *MongoInstance
	LOGS  *Collection
	TAGS  *Collection
	POSTS *Collection
)

var ok = Error{
	Code: 0,
	Msg:  "Ok",
}

func TestConnection(t *testing.T) {

	Convey("Should be able to create new mongo instance and connect to it", t, func() {
		Mg = NewMongoInstance(DB_NAME, DB_URL)
		//connect
		x := Mg.Connect()
		fmt.Println("mg=", Mg)
		So(x, ShouldEqual, ok)
		So(Mg.Connected, ShouldEqual, true)
		Convey("Should be able to get database list", func() {
			status := Mg.ListDatabaseNames()
			So(status, ShouldEqual, ok)
		})

	})
}

func TestCollections(t *testing.T) {
	//register and create index collections : logs, tags, posts
	Convey("Should be able to create index for collections :logs, tags, posts", t, func() {
		fmt.Println(Mg)

		LOGS = NewCollection("logs", Mg).SetIndex(log_index)
		status := LOGS.CreateIndex()
		So(status, ShouldEqual, ok)
		TAGS = NewCollection("tags", Mg).SetIndex(tag_index)
		status = TAGS.CreateIndex()
		So(status, ShouldEqual, ok)
		POSTS = NewCollection("posts", Mg).SetIndex(post_index)
		status = POSTS.CreateIndex()
		So(status, ShouldEqual, ok)

		Convey("Check mongo instance that have 3 registred collections", func() {

			So(status, ShouldEqual, ok)
		})

	})

}

// testing save
func TestSaveDocuments(t *testing.T) {

	Convey("Should be able to create 2 documents in :logs, tags, posts", t, func() {

		Convey("Save 2 documents in logs", func() {
			data := []Log{
				{
					Key:   "test1",
					Value: "value1",
				},
				{
					Key:   "test2",
					Value: "value2",
				},
			}

			for _, k := range data {
				x := NewDocument(LOGS).SetData(k).Save()
				So(x, ShouldEqual, ok)
			}

		})

		Convey("Save 2 documents in tags", func() {
			data := []Tag{
				{
					Value: "tag1",
					Count: 1,
				},
				{
					Value: "tag2",
					Count: 1,
				},
			}

			for _, k := range data {
				x := NewDocument(TAGS).SetData(k).Save()
				So(x, ShouldEqual, ok)
			}

		})

		Convey("Save 2 documents in posts", func() {
			data := []Post{
				{
					PostId: "p1",
					Author: "user1",
					Title:  "post title 1",
					Body:   "this body of post 1",
					Views:  1,
				},
				{
					PostId: "p2",
					Author: "user1",
					Title:  "post title 2",
					Body:   "this body of post 2",
					Views:  1,
				},
			}

			for _, k := range data {
				x := NewDocument(POSTS).SetData(k).Save()
				So(x, ShouldEqual, ok)
			}

		})

	})
}

// testing count
func TestCountDocuments(t *testing.T) {
	Convey("Should be able to count documents in :logs, tags, posts \n", t, func() {
		//count all logs:
		count, x := NewDocument(LOGS).CountAll()
		fmt.Println(count)
		So(x, ShouldEqual, ok)
		count, x = NewDocument(POSTS).CountAll()
		fmt.Println(count)
		So(x, ShouldEqual, ok)
		count, x = NewDocument(TAGS).CountAll()
		fmt.Println(count)
		So(x, ShouldEqual, ok)
	})
}

// testing checksaved
func TestCheckSavedDocuments(t *testing.T) {
	Convey("Should be able to check saved documents in logs, posts, tags \n", t, func() {
		Convey("Check Saved document in logs", func() {
			filter := Query{
				QType: QFilter,
				Kv:    KV{Key: "key", Value: "test1"},
			}
			x := NewDocument(LOGS).SetFilter(filter.Get()).CheckSaved()
			So(x, ShouldEqual, true)

		})

		//TODO for other collections is the same

	})
}

// testing update , incr
func TestUpdateDocuments(t *testing.T) {
	Convey("Should be able to update documents in logs, posts, tags \n", t, func() {
		Convey("update document key='test1' in logs", func() {
			filter := Query{
				QType: QFilter,
				Kv:    KV{Key: "key", Value: "test1"},
			}
			data := Query{
				QType: QSet,
				Data: bson.D{
					{Key: "value", Value: "test1updated"},
				},
			}
			x := NewDocument(LOGS).SetFilter(filter.Get()).SetUpdate(data.Get()).Update()
			So(x, ShouldEqual, ok)

		})
		Convey("update title document post_id='p1' in posts", func() {
			filter := Query{
				QType: QFilter,
				Kv:    KV{Key: "post_id", Value: "p1"},
			}
			data := Query{
				QType: QSet,
				Data: bson.D{
					{Key: "title", Value: "post 1 title updated"},
				},
			}
			x := NewDocument(POSTS).SetFilter(filter.Get()).SetUpdate(data.Get()).Update()
			So(x, ShouldEqual, ok)

			//incr views
			x = NewDocument(POSTS).SetFilter(filter.Get()).Incr("views", "1")
			So(x, ShouldEqual, ok)

		})

	})
}

// testing get, getcursor, paginate
func TestGetDocuments(t *testing.T) {
	Convey("Should be able to get saved documents in logs, posts, tags \n", t, func() {
		Convey("Get document by key=test2 in logs", func() {
			filter := Query{
				QType: QFilter,
				Kv:    KV{Key: "key", Value: "test2"},
			}
			var log Log
			err := NewDocument(LOGS).SetFilter(filter.Get()).Get().Decode(&log)
			So(err, ShouldEqual, nil)
			So(log.Key, ShouldEqual, "test2")
			So(log.Value, ShouldEqual, "value2")

		})

		//TODO for other collections is the same

	})
}

// testing delete
func TestDeleteDocuments(t *testing.T) {
	Convey("Should be able to get delte document in logs, posts, tags \n", t, func() {
		Convey("Delete document in tag, value=tag1", func() {
			filter := Query{
				QType: QFilter,
				Kv:    KV{Key: "value", Value: "tag1"},
			}

			x := NewDocument(TAGS).SetFilter(filter.Get()).Delete()
			So(x, ShouldEqual, ok)

		})

		//TODO for other collections is the same

	})
}

// testing paginate
func TestPaginateDocuments(t *testing.T) {
	Convey("Should be able to get paginate documents in posts \n", t, func() {

		posts := []Post{}
		pagination, x := NewDocument(POSTS).Paginate(10, 1, &posts)
		fmt.Println(posts)

		So(pagination.Total, ShouldEqual, 2)
		for _, k := range posts {
			fmt.Println(k)
		}
		So(x, ShouldEqual, ok)
	})
}
