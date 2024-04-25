package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOption func(*options.FindOptions)

func FindOptionWithSkip(skip int64) FindOption {
	return func(fo *options.FindOptions) {
		fo.SetSkip(skip)
	}
}

func FindOptionWithLimit(limit int64) FindOption {
	return func(fo *options.FindOptions) {
		fo.SetLimit(limit)
	}
}

func FindOptionWithSort(sort bson.D) FindOption {
	return func(fo *options.FindOptions) {
		if len(sort) > 0 {
			fo.SetSort(sort)
		}
	}
}

func FindOptionWithPage(pageIndex int64, pageSize int64) FindOption {
	return func(fo *options.FindOptions) {
		fo.SetLimit(pageSize)
		if pageIndex < 1 {
			pageIndex = 1
		}
		fo.SetSkip(pageSize * (pageIndex - 1))
	}
}

// result with specified fields
func FindOptionWithSpecifiedFields(fieldNameList []string) FindOption {
	return func(fo *options.FindOptions) {
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

func FindOptionWithFieldSort(fieldName string, isAsc bool) FindOption {
	return func(fo *options.FindOptions) {
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
