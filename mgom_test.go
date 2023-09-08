package mgom

import (
	"fmt"
	"testing"
)

// simple example test
// TODO all testing cases with goconvey

const (
	DB_NAME = "mgom_te"
	DB_URL  = "mongodb://localhost:27017/"
)

type Log struct {
	Key   string
	Value string
}

var log_indexs = []IndexFactory{
	{
		T: IdxUNIQUE,
		Values: map[string]string{
			"key": "true",
		},
	},
}

func TestMgom(t *testing.T) {

	//Mongo instance should be global pointer variable
	var Mg *MongoInstance
	Mg = NewMongoInstance(DB_NAME, DB_URL)
	fmt.Println(Mg)
	//connect
	code := Mg.Connect()
	//test get database list
	Mg.ListDatabaseNames()

	if Mg.Connected {
		fmt.Println("connected:", code)
		fmt.Println(Mg)

		LOGS := NewCollection("logs", Mg).SetIndex(log_indexs)
		LOGS.CreateIndex()

		//insert new doc to logs
		data := Log{
			Key:   "testing-key",
			Value: "hello-value",
		}

		x := NewDocument(LOGS).SetData(data).Save()
		fmt.Println(x)

	}

}
