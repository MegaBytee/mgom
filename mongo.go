package mgom

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Version = "0.0.2"

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	DbName    string
	BaseUrl   string
	Uri       string
	Context   context.Context
	Client    *mongo.Client
	Db        *mongo.Database
	Connected bool
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

func (mg *MongoInstance) GetCollection(value string) *mongo.Collection {
	return mg.Db.Collection(value)

}

func NewMongoInstance(dbName, baseUrl string) *MongoInstance {
	mg := &MongoInstance{
		Connected: false,
	}
	return mg.SetDbName(dbName).SetBaseUrl(baseUrl).SetUri()
}

func (mg *MongoInstance) Connect() Error {

	var err error
	var cancel context.CancelFunc

	mg.Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mg.Uri))

	mg.Context, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mg.Db = mg.Client.Database(mg.DbName)
	//mg.ListDatabaseNames()
	mg.Connected = true

	return handleErrors(CONNECT, err)

}

func (mg *MongoInstance) ListDatabaseNames() Error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	databases, err := mg.Client.ListDatabaseNames(ctx, bson.M{})
	fmt.Println("databases:", databases)
	return handleErrors(DEFAULT, err)
}
