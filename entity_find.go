package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IEntityFind interface {
	CountByFilter(filter interface{}, opts ...MongodbrCountOption) (count int64, err error)
	CountAll(opts ...WithContextOptions) (count int64, err error)

	// find
	FindAll(list interface{}, opts ...MongodbrFindOption) error
	FindListByFilter(filter interface{}, list interface{}, opts ...MongodbrFindOption) error
	FindListResultByFilter(filter interface{}, opts ...MongodbrFindOption) IFindResult
	FindListByObjectIdList(idList []primitive.ObjectID, list interface{}, opts ...MongodbrFindOption) error

	FindOneByObjectId(id primitive.ObjectID, v interface{}, opts ...MongodbrFindOneOption) error
	FindOne(filter interface{}, v interface{}, opts ...MongodbrFindOneOption) error

	Distinct(fieldName string, filter interface{}, opts ...*WithContextOptions) ([]interface{}, error)
}

var _ IEntityFind = (*MongoCol)(nil)

// #region IEntityFind Members

func (r *MongoCol) CountByFilter(filter interface{}, opts ...MongodbrCountOption) (int64, error) {
	// handle options
	cOptions := &MongodbrCountOptions{
		CountOptions: options.Count(),
	}
	for _, eachOpt := range opts {
		eachOpt(cOptions)
	}
	// handle context
	ctx, cancel := CreateContextAndCancelWith(r.configuration, cOptions.WithCtx)
	defer cancel()

	total, err := r.collection.CountDocuments(ctx, filter, cOptions.CountOptions)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *MongoCol) CountAll(opts ...WithContextOptions) (count int64, err error) {
	// handle options
	cOptions := &WithContextOptions{}
	for _, eachOpt := range opts {
		cOptions.WithCtx = eachOpt.WithCtx
	}
	// handle context
	ctx, cancel := CreateContextAndCancelWith(r.configuration, cOptions.WithCtx)
	defer cancel()

	total, err := r.collection.EstimatedDocumentCount(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *MongoCol) FindAll(list interface{}, opts ...MongodbrFindOption) error {
	return r.FindListByFilter(bson.M{}, list, opts...)
}

// 根据条件来筛选
// v,集合值,
func (r *MongoCol) FindListByFilter(filter interface{}, list interface{}, opts ...MongodbrFindOption) error {
	//设置默认搜索参数
	findOptions := &MongodbrFindOptions{
		FindOptions: options.Find(),
	}
	for _, o := range opts {
		o(findOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, findOptions.WithCtx)
	defer cancel()

	if findOptions.Sort == nil {
		// if sort is nil,then set default sort with configuration
		if r.configuration.setDefaultSort != nil {
			r.configuration.setDefaultSort(findOptions.FindOptions)
		}
	}

	cur, err := r.collection.Find(ctx, filter, findOptions.FindOptions)
	if err != nil {
		return err
	}
	result := &findResult{
		context:       ctx,
		configuration: r.configuration,
		cur:           cur,
	}
	err = result.All(list)
	if err != nil {
		return err
	}
	return nil
}

// 根据条件来筛选
// v,集合值,
func (r *MongoCol) FindListResultByFilter(filter interface{}, opts ...MongodbrFindOption) IFindResult {
	//设置默认搜索参数
	findOptions := &MongodbrFindOptions{
		FindOptions: options.Find(),
	}
	for _, o := range opts {
		o(findOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, findOptions.WithCtx)
	defer cancel()

	if findOptions.Sort == nil {
		// if sort is nil,then set default sort with configuration
		if r.configuration.setDefaultSort != nil {
			r.configuration.setDefaultSort(findOptions.FindOptions)
		}
	}

	cur, err := r.collection.Find(ctx, filter, findOptions.FindOptions)
	if err != nil {
		return &findResult{
			context:       ctx,
			configuration: r.configuration,
			err:           err,
		}
	}
	return &findResult{
		context:       ctx,
		configuration: r.configuration,
		cur:           cur,
	}
}

// 根据_id列表来查找，返回的是对象的指针
func (r *MongoCol) FindListByObjectIdList(idList []primitive.ObjectID, list interface{}, opts ...MongodbrFindOption) error {
	return r.FindListByFilter(bson.M{"_id": bson.M{
		"$in": idList,
	},
	}, list, opts...)
}

// 根据_id来查找，返回的是对象的指针
func (r *MongoCol) FindOneByObjectId(id primitive.ObjectID, v interface{}, opts ...MongodbrFindOneOption) error {
	return r.FindOne(bson.M{"_id": id}, v, opts...)
}

// 查找一条记录
func (r *MongoCol) FindOne(filter interface{}, v interface{}, opts ...MongodbrFindOneOption) error {
	//设置默认搜索参数
	mOptions := &MongodbrFindOneOptions{
		FindOneOptions: options.FindOne(),
	}
	for _, eachOpt := range opts {
		eachOpt(mOptions)
	}
	ctx, cancel := CreateContextAndCancelWith(r.configuration, mOptions.WithCtx)
	defer cancel()

	// find one
	res := r.collection.FindOne(ctx, filter, mOptions.FindOneOptions)
	err := res.Err()
	if err != nil {
		return err
	}
	result := &findResult{
		context:       ctx,
		configuration: r.configuration,
		res:           res,
	}
	err = result.One(v)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoCol) Distinct(fieldName string, filter interface{}, opts ...*WithContextOptions) ([]interface{}, error) {
	// handle options
	cOptions := NewWithContextOptions().MergeWithContextOptions(opts...)
	// handle context
	ctx, cancel := CreateContextAndCancelWith(r.configuration, cOptions.WithCtx)
	defer cancel()

	return r.collection.Distinct(ctx, fieldName, filter)
}

// #endregion
