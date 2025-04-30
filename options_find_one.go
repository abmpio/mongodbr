package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbrFindOneOptions struct {
	*options.FindOneOptions
	WithContextOptions
}

type MongodbrFindOneOption func(*MongodbrFindOneOptions)

func (o *MongodbrFindOneOptions) ensureFindOneOptionsInit() {
	if o.FindOneOptions == nil {
		o.FindOneOptions = options.FindOne()
	}
}

// merge MongodbrFindOneOption list and return one *MongodbrFindOneOptions
func MergeMongodbrFindOneOption(opts ...MongodbrFindOneOption) *MongodbrFindOneOptions {
	findOneOptions := &MongodbrFindOneOptions{
		FindOneOptions: options.FindOne(),
	}
	for _, eachOpt := range opts {
		eachOpt(findOneOptions)
	}
	return findOneOptions
}

// MongodbrFindOneOption with context
func MongodbrFindOneOptionWithContext(ctx context.Context) MongodbrFindOneOption {
	return func(mfoo *MongodbrFindOneOptions) {
		mfoo.WithCtx = ctx
	}
}

// MongodbrFindOneOption with specified fields
func MongodbrFindOneOptionWithSpecifiedFields(fieldNameList []string) MongodbrFindOneOption {
	return func(mfoo *MongodbrFindOneOptions) {
		if len(fieldNameList) <= 0 {
			return
		}
		mfoo.ensureFindOneOptionsInit()
		projection := bson.M{}
		for _, eachFieldName := range fieldNameList {
			projection[eachFieldName] = 1
		}
		mfoo.SetProjection(projection)
	}
}

// MongodbrFindOneOption with field sort
func MongodbrFindOneOptionWithFieldSort(fieldName string, isAsc bool) MongodbrFindOneOption {
	return func(mfoo *MongodbrFindOneOptions) {
		mfoo.ensureFindOneOptionsInit()
		if mfoo.Sort == nil {
			mfoo.Sort = bson.D{}
		}
		sortV, ok := mfoo.Sort.(bson.D)
		if !ok {
			return
		}
		sortValue := 1
		if !isAsc {
			sortValue = -1
		}
		sortV = append(sortV, bson.E{
			Key:   fieldName,
			Value: sortValue,
		})
		mfoo.SetSort(sortV)
	}
}
