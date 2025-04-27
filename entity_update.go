package mongodbr

import (
	"github.com/abmpio/mongodbr/builder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// update
type IEntityUpdate interface {
	FindOneAndUpdate(entity IEntity, opts ...MongodbrFindOneAndUpdateOption) error
	FindOneAndUpdateWithId(objectId primitive.ObjectID, update interface{}, opts ...MongodbrFindOneAndUpdateOption) error
	UpdateOne(filter interface{}, update interface{}, opts ...MongodbrUpdateOption) error
	UpdateMany(filter interface{}, update interface{}, opts ...MongodbrUpdateOption) (interface{}, error)
}

var _ IEntityUpdate = (*MongoCol)(nil)

// #region update members

func (r *MongoCol) FindOneAndUpdate(entity IEntity, opts ...MongodbrFindOneAndUpdateOption) error {
	objectId := entity.GetObjectId()
	update := builder.NewBsonBuilder().NewOrUpdateSet(entity).ToValue()
	return r.FindOneAndUpdateWithId(objectId, update, opts...)
}

func (r *MongoCol) FindOneAndUpdateWithId(objectId primitive.ObjectID, update interface{}, opts ...MongodbrFindOneAndUpdateOption) error {
	uOptions := options.FindOneAndUpdate()
	uOptions.SetUpsert(false)

	// handle options
	mongodbrUOptions := &MongodbrFindOneAndUpdateOptions{
		FindOneAndUpdateOptions: uOptions,
	}
	for _, eachOpt := range opts {
		eachOpt(mongodbrUOptions)
	}
	// handle context
	ctx, cancel := CreateContextWith(r.configuration, mongodbrUOptions.WithCtx)
	defer cancel()

	if err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectId},
		update,
		mongodbrUOptions.FindOneAndUpdateOptions,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (r *MongoCol) UpdateOne(filter interface{}, update interface{}, opts ...MongodbrUpdateOption) error {
	// handle options
	uOptions := &MongodbrUpdateOptions{
		UpdateOptions: options.Update(),
	}
	for _, eachOpt := range opts {
		eachOpt(uOptions)
	}
	// handle context
	ctx, cancel := CreateContextWith(r.configuration, uOptions.WithCtx)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, filter, update, uOptions.UpdateOptions)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoCol) UpdateMany(filter interface{}, update interface{}, opts ...MongodbrUpdateOption) (interface{}, error) {
	// handle options
	uOptions := &MongodbrUpdateOptions{
		UpdateOptions: options.Update(),
	}
	for _, eachOpt := range opts {
		eachOpt(uOptions)
	}
	// handle context
	ctx, cancel := CreateContextWith(r.configuration, uOptions.WithCtx)
	defer cancel()

	result, err := r.collection.UpdateMany(ctx, filter, update, uOptions.UpdateOptions)
	if err != nil {
		if result != nil {
			return result.UpsertedID, err
		} else {
			return nil, err
		}
	}

	if result != nil {
		return result.UpsertedID, nil
	}
	return nil, nil
}

// #endregion
