package mongodbr

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IEntityIndex interface {
	// index
	CreateIndex(indexModel mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error)
	CreateIndexes(indexModelList []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error)
	MustCreateIndex(indexModel mongo.IndexModel, opts ...*options.CreateIndexesOptions)
	MustCreateIndexes(indexModelList []mongo.IndexModel, opts ...*options.CreateIndexesOptions)
	DeleteIndex(name string) (err error)
	DeleteAllIndexes() (err error)
	ListIndexes() (indexes []map[string]interface{}, err error)
	ExistIndex(name string) (bool, error)
}

var _ IEntityIndex = (*MongoCol)(nil)

// #region indexes members

func (r *MongoCol) CreateIndex(indexModel mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	ctx, cancel := CreateContextAndCancel(r.configuration)
	defer cancel()
	name, err := r.collection.Indexes().CreateOne(ctx, indexModel, asOptionListers(opts)...)
	if err != nil {
		return "", err
	}
	return name, nil
}

func indexModelName(indexModel mongo.IndexModel) (string, bool, error) {
	if indexModel.Options == nil {
		return "", false, nil
	}
	indexOptions := &options.IndexOptions{}
	for _, applyOption := range indexModel.Options.List() {
		if err := applyOption(indexOptions); err != nil {
			return "", false, err
		}
	}
	if indexOptions.Name == nil {
		return "", false, nil
	}
	return *indexOptions.Name, true, nil
}

func (r *MongoCol) CreateIndexes(indexModelList []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	ctx, cancel := CreateContextAndCancel(r.configuration)
	defer cancel()

	notExistList := make([]mongo.IndexModel, 0)
	for _, eachIndexModel := range indexModelList {
		if indexName, ok, err := indexModelName(eachIndexModel); err != nil {
			return nil, err
		} else if ok {
			exist, err := r.ExistIndex(indexName)
			if err != nil {
				return nil, err
			}
			if exist {
				// exist ,continue
				continue
			}
		}
		notExistList = append(notExistList, eachIndexModel)
	}
	if len(notExistList) > 0 {
		return r.collection.Indexes().CreateMany(ctx, indexModelList, asOptionListers(opts)...)
	}
	return []string{}, nil
}

func (r *MongoCol) MustCreateIndex(indexModel mongo.IndexModel, opts ...*options.CreateIndexesOptions) {
	r.CreateIndex(indexModel, opts...)
}

func (r *MongoCol) MustCreateIndexes(indexModelList []mongo.IndexModel, opts ...*options.CreateIndexesOptions) {
	r.CreateIndexes(indexModelList, opts...)
}

func (r *MongoCol) DeleteIndex(name string) (err error) {
	ctx, cancel := CreateContextAndCancel(r.configuration)
	defer cancel()

	err = r.collection.Indexes().DropOne(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoCol) DeleteAllIndexes() (err error) {
	ctx, cancel := CreateContextAndCancel(r.configuration)
	defer cancel()

	err = r.collection.Indexes().DropAll(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoCol) ListIndexes() (indexes []map[string]interface{}, err error) {
	ctx, cancel := CreateContextAndCancel(r.configuration)
	defer cancel()

	cur, err := r.collection.Indexes().List(ctx)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &indexes); err != nil {
		return nil, err
	}
	return indexes, nil
}

func (r *MongoCol) ExistIndex(name string) (bool, error) {
	list, err := r.ListIndexes()
	if err != nil {
		return false, err
	}
	for _, eachIndex := range list {
		indexName, ok := eachIndex["name"]
		if ok && indexName == name {
			return true, nil
		}
	}
	return false, nil
}

// #endregion
