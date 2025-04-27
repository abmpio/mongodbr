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
	FindAll(opts ...MongodbrFindOption) IFindResult
	FindByObjectId(id primitive.ObjectID, opts ...MongodbrFindOneOption) IFindResult
	FindListByObjectIdList(idList []primitive.ObjectID, opts ...MongodbrFindOption) IFindResult
	FindOne(filter interface{}, opts ...MongodbrFindOneOption) IFindResult
	FindByFilter(filter interface{}, opts ...MongodbrFindOption) IFindResult

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
	ctx, cancel := CreateContextWith(r.configuration, cOptions.WithCtx)
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
	ctx, cancel := CreateContextWith(r.configuration, cOptions.WithCtx)
	defer cancel()

	total, err := r.collection.EstimatedDocumentCount(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *MongoCol) FindAll(opts ...MongodbrFindOption) IFindResult {
	return r.FindByFilter(bson.M{}, opts...)
}

// 根据_id来查找，返回的是对象的指针
func (r *MongoCol) FindByObjectId(id primitive.ObjectID, opts ...MongodbrFindOneOption) IFindResult {
	return r.FindOne(bson.M{"_id": id}, opts...)
}

// 根据_id列表来查找，返回的是对象的指针
func (r *MongoCol) FindListByObjectIdList(idList []primitive.ObjectID, opts ...MongodbrFindOption) IFindResult {
	return r.FindByFilter(bson.M{"_id": bson.M{
		"$in": idList,
	},
	}, opts...)
}

// 查找一条记录
func (r *MongoCol) FindOne(filter interface{}, opts ...MongodbrFindOneOption) IFindResult {
	//设置默认搜索参数
	mOptions := &MongodbrFindOneOptions{
		FindOneOptions: options.FindOne(),
	}
	for _, eachOpt := range opts {
		eachOpt(mOptions)
	}
	ctx, cancel := CreateContextWith(r.configuration, mOptions.WithCtx)
	defer cancel()

	res := r.collection.FindOne(ctx, filter, mOptions.FindOneOptions)
	if res.Err() != nil {
		return &findResult{
			context:       ctx,
			configuration: r.configuration,
			err:           res.Err(),
		}
	}
	return &findResult{
		context:       ctx,
		configuration: r.configuration,
		res:           res,
	}
}

// 根据条件来筛选
func (r *MongoCol) FindByFilter(filter interface{}, opts ...MongodbrFindOption) IFindResult {
	//设置默认搜索参数
	findOptions := &MongodbrFindOptions{
		FindOptions: options.Find(),
	}
	for _, o := range opts {
		o(findOptions)
	}
	ctx, cancel := CreateContextWith(r.configuration, findOptions.WithCtx)
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

func (r *MongoCol) Distinct(fieldName string, filter interface{}, opts ...*WithContextOptions) ([]interface{}, error) {
	// handle options
	cOptions := NewWithContextOptions().MergeWithContextOptions(opts...)
	// handle context
	ctx, cancel := CreateContextWith(r.configuration, cOptions.WithCtx)
	defer cancel()

	return r.collection.Distinct(ctx, fieldName, filter)
}

// #endregion
