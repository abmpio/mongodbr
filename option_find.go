package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbrFindOptions struct {
	*options.FindOptions
	WithContextOptions
}

type MongodbrFindOption func(*MongodbrFindOptions)

func (o *MongodbrFindOptions) ensureFindOptionsInit() {
	if o.FindOptions == nil {
		o.FindOptions = options.Find()
	}
}

// merge MongodbrFindOption list and return one *MongodbrFindOptions
func MergeMongodbrFindOption(opts ...MongodbrFindOption) *MongodbrFindOptions {
	o := &MongodbrFindOptions{
		FindOptions: options.Find(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

func MongodbrFindOptionWithSkip(skip int64) MongodbrFindOption {
	return func(fo *MongodbrFindOptions) {
		fo.ensureFindOptionsInit()
		fo.SetSkip(skip)
	}
}

func MongodbrFindOptionWithLimit(limit int64) MongodbrFindOption {
	return func(fo *MongodbrFindOptions) {
		fo.ensureFindOptionsInit()
		fo.SetLimit(limit)
	}
}

func MongodbrFindOptionWithSort(sort bson.D) MongodbrFindOption {
	return func(fo *MongodbrFindOptions) {
		fo.ensureFindOptionsInit()
		if len(sort) > 0 {
			fo.SetSort(sort)
		}
	}
}

func MongodbrFindOptionWithPage(pageIndex int64, pageSize int64) MongodbrFindOption {
	return func(fo *MongodbrFindOptions) {
		fo.ensureFindOptionsInit()
		fo.SetLimit(pageSize)
		if pageIndex < 1 {
			pageIndex = 1
		}
		fo.SetSkip(pageSize * (pageIndex - 1))
	}
}

// result with specified fields
func MongodbrFindOptionWithSpecifiedFields(fieldNameList []string) MongodbrFindOption {
	return func(fo *MongodbrFindOptions) {
		if len(fieldNameList) <= 0 {
			return
		}
		fo.ensureFindOptionsInit()
		projection := bson.M{}
		for _, eachFieldName := range fieldNameList {
			projection[eachFieldName] = 1
		}
		fo.SetProjection(projection)
	}
}

func MongodbrFindOptionWithFieldSort(fieldName string, isAsc bool) MongodbrFindOption {
	return func(fo *MongodbrFindOptions) {
		if fo.Sort == nil {
			fo.Sort = bson.D{}
		}
		sortV, ok := fo.Sort.(bson.D)
		if !ok {
			return
		}
		fo.ensureFindOptionsInit()
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
