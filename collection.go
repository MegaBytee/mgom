package mgom

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	Name  string
	Mg    *MongoInstance
	Index []IndexFactory
}

func (c *Collection) SetIndex(idx []IndexFactory) *Collection {
	c.Index = idx
	return c
}

func (c *Collection) GetCollection() *mongo.Collection {
	return c.Mg.GetCollection(c.Name)
}

func NewCollection(name string, mg *MongoInstance) *Collection {
	c := &Collection{
		Name: name,
		Mg:   mg,
	}
	return c
}

func (c *Collection) CreateIndex() Error {

	err := Error{}
	for v := range c.Index {
		err = c.Index[v].CreateIndex(c)
	}
	return err
}
