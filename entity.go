package mongodbr

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Entity struct {
	ObjectId bson.ObjectID `json:"objectId,omitempty" bson:"_id"`
}

// modify IEntity object
type EntityOption = func(e IEntity)

type IEntity interface {
	GetObjectId() bson.ObjectID
}

type IEntityBeforeCreate interface {
	BeforeCreate()
}

type IEntityBeforeUpdate interface {
	BeforeUpdate()
}

// 创建时设置对象的基本信息
func (entity *Entity) BeforeCreate() {
	if entity.ObjectId == bson.NilObjectID {
		entity.ObjectId = bson.NewObjectID()
	}
}

func (entity Entity) GetObjectId() bson.ObjectID {
	return entity.ObjectId
}
