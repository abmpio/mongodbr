package mongodbr

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoCol struct {
	configuration *Configuration
	collection    *mongo.Collection
}

// new MongoCol instance, panic if col is nil
func NewMongoCol(col *mongo.Collection, opts ...*Configuration) *MongoCol {
	if col == nil {
		panic(errors.New("col cannot be nil"))
	}
	c := NewConfiguration()
	if len(opts) > 0 {
		c = opts[0]
	}
	mongoCol := &MongoCol{
		configuration: c,
		collection:    col,
	}
	return mongoCol
}

// RepositoryBase represents a mongodb repository
type RepositoryBase struct {
	documentName string
	*MongoCol
}

var _ IRepository = (*RepositoryBase)(nil)

// new一个新的实例
func NewRepositoryBase(getDbCollection func() *mongo.Collection, opts ...RepositoryOption) (*RepositoryBase, error) {
	if getDbCollection == nil {
		err := fmt.Errorf("getDbCollection参数不能为nil")
		return nil, err
	}
	coll := getDbCollection()
	repository := &RepositoryBase{
		MongoCol:     NewMongoCol(coll),
		documentName: coll.Name(),
	}
	for _, eachItem := range opts {
		eachItem(repository.configuration)
	}
	return repository, nil
}

// #region create members

func (r *RepositoryBase) Create(item interface{}, opts ...MongodbrInsertOneOption) (id bson.ObjectID, err error) {
	if item == nil {
		return bson.NilObjectID, fmt.Errorf("item is nil,col:%s", r.documentName)
	}

	insertOneOptions := &MongodbrInsertOneOptions{
		InsertOneOptions: &options.InsertOneOptions{},
	}
	for _, o := range opts {
		o(insertOneOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, insertOneOptions.WithCtx)
	defer cancel()

	r.onBeforeCreate(item)
	res, err := r.collection.InsertOne(ctx, item, insertOneOptions)
	if err != nil {
		return bson.NilObjectID, err
	}
	if id, ok := res.InsertedID.(bson.ObjectID); ok {
		return id, nil
	}
	return bson.NilObjectID, ErrInvalidType
}

func (r *RepositoryBase) CreateMany(itemList []interface{}, opts ...MongodbrInsertManyOption) (ids []bson.ObjectID, err error) {
	if len(itemList) <= 0 {
		return nil, nil
	}

	insertManyOptions := &MongodbrInsertManyOptions{
		InsertManyOptions: &options.InsertManyOptions{},
	}
	for _, o := range opts {
		o(insertManyOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, insertManyOptions.WithCtx)
	defer cancel()

	for index := range itemList {
		r.onBeforeCreate(itemList[index])
	}
	res, err := r.collection.InsertMany(ctx, itemList, insertManyOptions)
	if err != nil {
		return nil, err
	}
	for _, v := range res.InsertedIDs {
		switch v := v.(type) {
		case bson.ObjectID:
			ids = append(ids, v)
		default:
			return nil, ErrInvalidType
		}
	}
	return ids, nil
}

// #endregion

func (r *RepositoryBase) ReplaceById(id bson.ObjectID, doc interface{}, opts ...MongodbrReplaceOption) (err error) {
	return r.Replace(bson.M{"_id": id}, doc, opts...)
}

func (r *RepositoryBase) Replace(filter interface{}, doc interface{}, opts ...MongodbrReplaceOption) (err error) {
	rOptions := &MongodbrReplaceOptions{
		ReplaceOptions: &options.ReplaceOptions{},
	}
	for _, o := range opts {
		o(rOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, rOptions.WithCtx)
	defer cancel()

	_, err = r.collection.ReplaceOne(ctx, filter, doc, rOptions)
	if err != nil {
		return err
	}
	return nil
}

// 删除指定id的记录
func (r *RepositoryBase) DeleteOne(id bson.ObjectID, opts ...MongodbrDeleteOption) (*mongo.DeleteResult, error) {
	deleteOptions := &MongodbrDeleteOptions{
		DeleteOneOptions:  &options.DeleteOneOptions{},
		DeleteManyOptions: &options.DeleteManyOptions{},
	}
	for _, o := range opts {
		o(deleteOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, deleteOptions.WithCtx)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id}, asOptionLister(deleteOptions.DeleteOneOptions))
	if err != nil {
		return result, err
	}

	return result, nil
}

// 删除指定条件的一条记录
func (r *RepositoryBase) DeleteOneByFilter(filter interface{}, opts ...MongodbrDeleteOption) (*mongo.DeleteResult, error) {
	deleteOptions := &MongodbrDeleteOptions{
		DeleteOneOptions:  &options.DeleteOneOptions{},
		DeleteManyOptions: &options.DeleteManyOptions{},
	}
	for _, o := range opts {
		o(deleteOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, deleteOptions.WithCtx)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, filter, asOptionLister(deleteOptions.DeleteOneOptions))
	if err != nil {
		return result, err
	}

	return result, nil
}

// 删除多条记录
func (r *RepositoryBase) DeleteMany(filter interface{}, opts ...MongodbrDeleteOption) (*mongo.DeleteResult, error) {
	if filter == nil {
		err := fmt.Errorf("无法删除多条%s记录,filter参数不能为null", r.documentName)
		return nil, err
	}
	deleteOptions := &MongodbrDeleteOptions{
		DeleteOneOptions:  &options.DeleteOneOptions{},
		DeleteManyOptions: &options.DeleteManyOptions{},
	}
	for _, o := range opts {
		o(deleteOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, deleteOptions.WithCtx)
	defer cancel()

	result, err := r.collection.DeleteMany(ctx, filter, asOptionLister(deleteOptions.DeleteManyOptions))
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *RepositoryBase) GetName() (name string) {
	return r.documentName
}

func (r *RepositoryBase) GetCollection() (c *mongo.Collection) {
	return r.collection
}

func (r *RepositoryBase) onBeforeCreate(item interface{}) {
	entityHookable, ok := item.(IEntityBeforeCreate)
	if !ok {
		return
	}
	entityHookable.BeforeCreate()
}
