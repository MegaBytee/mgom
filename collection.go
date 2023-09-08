package mgom

import (
	"fmt"

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
	fmt.Print("GetCollection=", c.Name)
	return c.Mg.GetCollection(c.Name)
}

func NewCollection(name string, mg *MongoInstance) *Collection {
	c := &Collection{
		Name: name,
		Mg:   mg,
	}
	return c
}

func (c *Collection) CreateIndex() int {

	code := 0
	for v := range c.Index {
		code = c.Index[v].CreateIndex(c)
	}
	return code
}
