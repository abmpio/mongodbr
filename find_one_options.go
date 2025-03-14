package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneOption func(*options.FindOneOptions)

// result with specified fields
func FindOneOptionWithSpecifiedFields(fieldNameList []string) FindOneOption {
	return func(fo *options.FindOneOptions) {
		if len(fieldNameList) <= 0 {
			return
		}
		projection := bson.M{}
		for _, eachFieldName := range fieldNameList {
			projection[eachFieldName] = 1
		}
		fo.SetProjection(projection)
	}
}

func FindOneOptionWithFieldSort(fieldName string, isAsc bool) FindOneOption {
	return func(fo *options.FindOneOptions) {
		if fo.Sort == nil {
			fo.Sort = bson.D{}
		}
		sortV, ok := fo.Sort.(bson.D)
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
		fo.SetSort(sortV)
	}
}
