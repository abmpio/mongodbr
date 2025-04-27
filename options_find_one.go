package mongodbr

import (
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

// merge contextOptions
func FindOneOptionsWithContextOptions(opts ...WithContextOptions) MongodbrFindOneOption {
	return func(mfoo *MongodbrFindOneOptions) {
		for _, eachCtx := range opts {
			mfoo.WithCtx = eachCtx.WithCtx
		}
	}
}

// result with specified fields
func FindOneOptionsWithSpecifiedFields(fieldNameList []string) MongodbrFindOneOption {
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

func FindOneOptionsWithFieldSort(fieldName string, isAsc bool) MongodbrFindOneOption {
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
