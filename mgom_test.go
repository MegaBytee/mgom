package mgom

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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

func TestConnection(t *testing.T) {

	Convey("Should be able to create new mongo instance and connect to it", t, func() {
		Mg = NewMongoInstance(DB_NAME, DB_URL)
		//connect
		Mg.Connect()
		So(Mg.Connected, ShouldEqual, true)
		Convey("Should be able to get database list", func() {
			status := Mg.ListDatabaseNames()
			So(status, ShouldEqual, 0)
		})

	})
}

func TestCollections(t *testing.T) {
	//register and create index collections : logs, tags, posts
	Convey("Should be able to create index for collections :logs, tags, posts", t, func() {
		fmt.Println(Mg)

		LOGS = NewCollection("logs", Mg).SetIndex(log_index)
		code := LOGS.CreateIndex()
		So(code, ShouldEqual, 0)
		TAGS = NewCollection("tags", Mg).SetIndex(tag_index)
		code = TAGS.CreateIndex()
		So(code, ShouldEqual, 0)
		POSTS = NewCollection("posts", Mg).SetIndex(post_index)
		code = POSTS.CreateIndex()
		So(code, ShouldEqual, 0)

		Convey("Check mongo instance that have 3 registred collections", func() {

			So(*Mg.Collections, ShouldEqual, 3)
		})

	})

}

// testing save, count, checksaved
func TestSaveDocuments(t *testing.T) {

}

// testing update , incr
func TestUpdateDocuments(t *testing.T) {

}

// testing get, getcursor, paginate
func TestGetDocuments(t *testing.T) {

}

// testing delete
func TestDeleteDocuments(t *testing.T) {

}
