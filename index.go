package mgom

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IndexType int

// enums
const (
	IdxUNIQUE IndexType = iota
	IdxTEXT
)

type IndexUnique struct {
	mIndex mongo.IndexModel
	key    string
	value  bool
}

func (x *IndexUnique) Set() {
	x.mIndex = mongo.IndexModel{
		Keys:    bson.M{x.key: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(x.value),
	}

}

type IndexText struct {
	mIndex mongo.IndexModel
	key    string
	weight int
}

func (x *IndexText) Set() {
	x.mIndex = mongo.IndexModel{
		Keys:    bson.D{{Key: x.key, Value: "text"}},
		Options: options.Index().SetWeights(bson.D{{Key: x.key, Value: x.weight}}),
	}

}

type IndexFactory struct {
	T      IndexType
	Values map[string]string
}

func (x *IndexFactory) CreateIndex(c *Collection) Error {
	err := Error{
		Code: -1,
		Msg:  "init",
	}
	switch x.T {
	case IdxUNIQUE:
		//fmt.Println("im here, create_index, indexfactory, unique", c)
		for v := range x.Values {
			idx := IndexUnique{
				key:   v,
				value: StringToBool(x.Values[v]),
			}
			idx.Set()
			//fmt.Println(" unique index, idx=", idx)
			err = createIndex(idx.mIndex, c)
		}

	case IdxTEXT:
		//fmt.Println("im here, create_index, indexfactory, text", c)
		for v := range x.Values {

			idx := IndexText{
				key:    v,
				weight: StringToInt(x.Values[v]),
			}

			idx.Set()
			//fmt.Println(" text index, idx=", idx)
			err = createIndex(idx.mIndex, c)
		}

	}

	return err

}

func createIndex(x mongo.IndexModel, c *Collection) Error {
	//fmt.Print("createIndex c=>", c)
	//fmt.Print("createIndex x=>", x)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error
	if c.Mg.Connected {
		_, err = c.GetCollection().Indexes().CreateOne(ctx, x)
	}
	return handleErrors(INDEX_CREATE, err)
}
