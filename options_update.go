package mongodbr

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOneAndUpdateOptions with context
type MongodbrFindOneAndUpdateOptions struct {
	*options.FindOneAndUpdateOptions
	WithContextOptions
}

type MongodbrFindOneAndUpdateOption func(*MongodbrFindOneAndUpdateOptions)

// merge MongodbrFindOneAndUpdateOption list and return one *MongodbrFindOneAndUpdateOptions
func MergeMongodbrFindOneAndUpdateOption(opts ...MongodbrFindOneAndUpdateOption) *MongodbrFindOneAndUpdateOptions {
	o := &MongodbrFindOneAndUpdateOptions{
		FindOneAndUpdateOptions: options.FindOneAndUpdate(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// UpdateOptions with context
type MongodbrUpdateOptions struct {
	*options.UpdateOptions
	WithContextOptions
}

type MongodbrUpdateOption func(*MongodbrUpdateOptions)

// merge MongodbrUpdateOption list and return one *MongodbrUpdateOption
func MergeMongodbrUpdateOption(opts ...MongodbrUpdateOption) *MongodbrUpdateOptions {
	o := &MongodbrUpdateOptions{
		UpdateOptions: options.Update(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// ReplaceOptions with context
type MongodbrReplaceOptions struct {
	*options.ReplaceOptions
	WithContextOptions
}

type MongodbrReplaceOption func(*MongodbrReplaceOptions)

// merge MongodbrReplaceOption list and return one *MongodbrReplaceOptions
func MergeMongodbrReplaceOption(opts ...MongodbrReplaceOption) *MongodbrReplaceOptions {
	o := &MongodbrReplaceOptions{
		ReplaceOptions: options.Replace(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// InsertOneOptions with context
type MongodbrInsertOneOptions struct {
	*options.InsertOneOptions
	WithContextOptions
}

type MongodbrInsertOneOption func(*MongodbrInsertOneOptions)

// merge MongodbrInsertOneOption list and return one *MongodbrInsertOneOptions
func MergeMongodbrInsertOneOption(opts ...MongodbrInsertOneOption) *MongodbrInsertOneOptions {
	o := &MongodbrInsertOneOptions{
		InsertOneOptions: options.InsertOne(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// InsertOneOptions with context
type MongodbrInsertManyOptions struct {
	*options.InsertManyOptions
	WithContextOptions
}

type MongodbrInsertManyOption func(*MongodbrInsertManyOptions)

// merge MongodbrInsertManyOption list and return one *MongodbrInsertManyOptions
func MergeMongodbrInsertManyOption(opts ...MongodbrInsertManyOption) *MongodbrInsertManyOptions {
	o := &MongodbrInsertManyOptions{
		InsertManyOptions: options.InsertMany(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// DeleteOptions with context
type MongodbrDeleteOptions struct {
	*options.DeleteOptions
	WithContextOptions
}

type MongodbrDeleteOption func(*MongodbrDeleteOptions)

// merge MongodbrDeleteOption list and return one *MongodbrDeleteOptions
func MergeMongodbrDeleteOption(opts ...MongodbrDeleteOption) *MongodbrDeleteOptions {
	o := &MongodbrDeleteOptions{
		DeleteOptions: options.Delete(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}
