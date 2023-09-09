package mgom

//version : v0.0.1
import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	DbName      string
	BaseUrl     string
	Uri         string
	Context     context.Context
	Client      *mongo.Client
	Db          *mongo.Database
	Collections *int
	Connected   bool
}

func (mg *MongoInstance) SetDbName(value string) *MongoInstance {
	mg.DbName = value
	return mg
}

func (mg *MongoInstance) SetBaseUrl(value string) *MongoInstance {
	mg.BaseUrl = value
	return mg
}
func (mg *MongoInstance) SetUri() *MongoInstance {
	mg.Uri = mg.BaseUrl + mg.DbName
	return mg
}
func (mg *MongoInstance) SetCollection(c *Collection) *MongoInstance {
	if mg.Collections == nil {
		mg.Collections = new(int)
		*mg.Collections = 0
	}
	*mg.Collections++
	return mg

}
func (mg *MongoInstance) GetCollection(value string) *mongo.Collection {
	return mg.Db.Collection(value)

}

func NewMongoInstance(dbName, baseUrl string) *MongoInstance {
	mg := &MongoInstance{
		Connected: false,
	}
	return mg.SetDbName(dbName).SetBaseUrl(baseUrl).SetUri()
}

func (mg *MongoInstance) Connect() int {

	var err error
	var cancel context.CancelFunc

	mg.Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mg.Uri))

	mg.Context, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mg.Db = mg.Client.Database(mg.DbName)
	//mg.ListDatabaseNames()
	code := handleErrors(CONNECT, err)
	if code == 0 {
		//connected
		mg.Connected = true
	}
	return code
}

func (mg *MongoInstance) ListDatabaseNames() int {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	databases, err := mg.Client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println("databases:", databases)
	return handleErrors(DEFAULT, err)
}
