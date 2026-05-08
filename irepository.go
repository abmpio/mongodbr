package mongodbr

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// 一个抽象的用来处理任意类型的mongodb的仓储基类
type IRepository interface {
	IEntityFind
	IEntityCreate
	IEntityUpdate
	IEntityDelete
	IEntityIndex
	IEntityBulkWrite

	// aggregate
	Aggregate(pipeline interface{}, dataList interface{}, opts ...MongodbrAggregateOption) (err error)

	// replace*
	ReplaceById(id bson.ObjectID, doc interface{}, opts ...MongodbrReplaceOption) (err error)
	Replace(filter interface{}, doc interface{}, opts ...MongodbrReplaceOption) (err error)

	GetName() (name string)
	GetCollection() (c *mongo.Collection)
}

type IEntityCreate interface {
	// create
	Create(data interface{}, opts ...MongodbrInsertOneOption) (id bson.ObjectID, err error)
	CreateMany(itemList []interface{}, opts ...MongodbrInsertManyOption) (ids []bson.ObjectID, err error)
}

type IEntityDelete interface {
	// delete
	DeleteOne(id bson.ObjectID, opts ...MongodbrDeleteOption) (*mongo.DeleteResult, error)
	DeleteOneByFilter(filter interface{}, opts ...MongodbrDeleteOption) (*mongo.DeleteResult, error)
	DeleteMany(filter interface{}, opts ...MongodbrDeleteOption) (*mongo.DeleteResult, error)
}
