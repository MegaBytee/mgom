package mgom

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QType int

const (
	QFilter QType = iota
	QSet
	QSearch
	QInc
)

type KV struct {
	Key   string
	Value string
}

type Query struct {
	QType QType
	Data  primitive.D
	Kv    KV
}

func (q *Query) Get() primitive.D {
	x := bson.D{}
	switch q.QType {
	case QFilter:
		x = bson.D{{Key: q.Kv.Key, Value: q.Kv.Value}}
		return x
	case QSet:
		x = bson.D{{Key: _SET, Value: q.Data}}
		return x
	case QSearch:
		x = bson.D{{Key: _TEXT, Value: bson.D{{Key: _SEARCH, Value: q.Kv.Value}}}}
		return x
	case QInc:
		x = bson.D{{Key: _INC, Value: bson.D{{Key: q.Kv.Key, Value: StringToInt(q.Kv.Value)}}}}
		return x

	default:
		return x

	}
}

/*
data := bson.D{
	{Key: "$set",
		Value: bson.D{
			{Key: key, Value: value},
			{Key: "comments", Value: x.Comments},
		},
	},
}
fields := bson.D{{Key: "tags", Value: 0}}
filter := bson.D{{Key: "hash", Value: x.Hash}}
filter := bson.D{
		{Key: "$text", Value: bson.D{{Key: "$search", Value: keyword}}},
	}
*/
