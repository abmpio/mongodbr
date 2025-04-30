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

// MongodbrFindOneAndUpdateOption with context
func MongodbrFindOneAndUpdateOptionWithContext(opts ...WithContextOptions) MongodbrFindOneAndUpdateOption {
	return func(mfoo *MongodbrFindOneAndUpdateOptions) {
		for _, eachCtx := range opts {
			mfoo.WithCtx = eachCtx.WithCtx
		}
	}
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

// MongodbrUpdateOption with context
func MongodbrUpdateOptionWithContext(opts ...WithContextOptions) MongodbrUpdateOption {
	return func(mfoo *MongodbrUpdateOptions) {
		for _, eachCtx := range opts {
			mfoo.WithCtx = eachCtx.WithCtx
		}
	}
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

// MongodbrReplaceOption with context
func MongodbrReplaceOptionWithContext(opts ...WithContextOptions) MongodbrReplaceOption {
	return func(mfoo *MongodbrReplaceOptions) {
		for _, eachCtx := range opts {
			mfoo.WithCtx = eachCtx.WithCtx
		}
	}
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

// MongodbrInsertOneOption with context
func MongodbrInsertOneOptionWithContext(opts ...WithContextOptions) MongodbrInsertOneOption {
	return func(mfoo *MongodbrInsertOneOptions) {
		for _, eachCtx := range opts {
			mfoo.WithCtx = eachCtx.WithCtx
		}
	}
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

// MongodbrInsertManyOption with context
func MongodbrInsertManyOptionWithContext(opts ...WithContextOptions) MongodbrInsertManyOption {
	return func(mfoo *MongodbrInsertManyOptions) {
		for _, eachCtx := range opts {
			mfoo.WithCtx = eachCtx.WithCtx
		}
	}
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

// MongodbrDeleteOption with context
func MongodbrDeleteOptionWithContext(opts ...WithContextOptions) MongodbrDeleteOption {
	return func(mfoo *MongodbrDeleteOptions) {
		for _, eachCtx := range opts {
			mfoo.WithCtx = eachCtx.WithCtx
		}
	}
}
