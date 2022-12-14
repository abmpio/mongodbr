package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Entity struct {
	ObjectId primitive.ObjectID `json:"objectId,omitempty" bson:"_id"`
}

type IEntity interface {
	GetObjectId() primitive.ObjectID
}

type IEntityBeforeCreate interface {
	BeforeCreate()
}

type IEntityBeforeUpdate interface {
	BeforeUpdate()
}

// 创建时设置对象的基本信息
func (entity *Entity) BeforeCreate() {
	if entity.ObjectId == primitive.NilObjectID {
		entity.ObjectId = primitive.NewObjectID()
	}
}

func (entity *Entity) GetObjectId() primitive.ObjectID {
	return entity.ObjectId
}

type FindOption func(*options.FindOptions)

// AggregateOptions handler pipeline
type AggregateOption func(*options.AggregateOptions)

type FindOneOption func(*options.FindOneOptions)
