package mgom

import (
	"context"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Document struct {
	collection *mongo.Collection
	fields     primitive.D
	filter     primitive.D
	data       any
	dUpdate    primitive.D
}

func (d *Document) SetFields(f primitive.D) *Document {
	if f == nil {
		d.fields = bson.D{}
		return d
	}
	d.fields = f
	return d
}

func (d *Document) SetFilter(f primitive.D) *Document {
	if f == nil {
		d.filter = bson.D{}
		return d
	}
	d.filter = f
	return d
}
func (d *Document) SetData(a any) *Document {
	d.data = a
	return d
}
func (d *Document) SetUpdate(a primitive.D) *Document {
	d.dUpdate = a
	return d
}

func (d *Document) SetCollection(c *Collection) *Document {
	d.collection = c.GetCollection()
	return d
}

func NewDocument(c *Collection) *Document {
	d := &Document{}
	return d.SetCollection(c).SetFields(nil).SetFilter(nil)

}

func (d *Document) CountAll() (int64, int) {
	opts := options.Count().SetHint("_id_")
	count, err := d.collection.CountDocuments(context.Background(), bson.D{}, opts)

	return count, handleErrors(COUNT, err)
}

func (d *Document) GetCursor(limit int64) (*mongo.Cursor, int) {
	opts := options.Find().SetLimit(limit)
	cursor, err := d.collection.Find(context.TODO(), d.filter, opts)
	return cursor, handleErrors(GET_CURSOR, err)
}

func (d *Document) PaginateWithSelect(limit, page int64) paginate.PagingQuery {

	return paginate.New(d.collection).Context(context.TODO()).Limit(limit).Page(page).Select(d.fields).Filter(d.filter)

}

func (d *Document) Get() *mongo.SingleResult {
	opts := options.FindOne().SetProjection(d.fields)
	return d.collection.FindOne(context.TODO(), d.filter, opts)
}

func (d *Document) Paginate(limit, page int64, x any) paginate.PaginationData {

	paginatedData, err := d.PaginateWithSelect(limit, page).Decode(x).Find()
	handleErrors(PAGINATE, err)
	return paginatedData.Pagination

}

// check if docs already saved in database
func (d *Document) CheckSaved() bool {
	var a any

	err := d.Get().Decode(&a)
	if err != nil && err == mongo.ErrNoDocuments {
		return false
	}
	return true
}

func (d *Document) Save() int {
	_, err := d.collection.InsertOne(context.TODO(), d.data)
	return handleErrors(SAVE, err)
}

func (d *Document) Update() int {
	_, err := d.collection.UpdateOne(context.TODO(), d.filter, d.dUpdate)
	return handleErrors(UPDATE, err)
}

func (d *Document) Delete() int {
	_, err := d.collection.DeleteOne(context.TODO(), d.filter)
	return handleErrors(DELETE, err)
}

func (d *Document) Incr(key string, value string) int {
	q := Query{
		QType: QInc,
		Kv:    KV{Key: key, Value: value},
	}

	d.dUpdate = q.Get()
	return d.Update()
}
