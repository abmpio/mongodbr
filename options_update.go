package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOneAndUpdateOptions with context
type MongodbrFindOneAndUpdateOptions struct {
	*options.FindOneAndUpdateOptions
	WithContextOptions
}

type MongodbrFindOneAndUpdateOption func(*MongodbrFindOneAndUpdateOptions)

// #region MongodbrFindOneAndUpdateOption members

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
func MongodbrFindOneAndUpdateOptionWithContext(ctx context.Context) MongodbrFindOneAndUpdateOption {
	return func(mfoo *MongodbrFindOneAndUpdateOptions) {
		mfoo.WithCtx = ctx
	}
}

// #endregion

// UpdateOptions with context
type MongodbrUpdateOptions struct {
	*options.UpdateOptions
	WithContextOptions
}

type MongodbrUpdateOption func(*MongodbrUpdateOptions)

// #region MongodbrUpdateOption members

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
func MongodbrUpdateOptionWithContext(ctx context.Context) MongodbrUpdateOption {
	return func(mfoo *MongodbrUpdateOptions) {
		mfoo.WithCtx = ctx
	}
}

// #endregion

// ReplaceOptions with context
type MongodbrReplaceOptions struct {
	*options.ReplaceOptions
	WithContextOptions
}

type MongodbrReplaceOption func(*MongodbrReplaceOptions)

// #region MongodbrReplaceOption Members

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
func MongodbrReplaceOptionWithContext(ctx context.Context) MongodbrReplaceOption {
	return func(mfoo *MongodbrReplaceOptions) {
		mfoo.WithCtx = ctx
	}
}

// #endregion

// InsertOneOptions with context
type MongodbrInsertOneOptions struct {
	*options.InsertOneOptions
	WithContextOptions
}

type MongodbrInsertOneOption func(*MongodbrInsertOneOptions)

// #region MongodbrInsertOneOption members

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
func MongodbrInsertOneOptionWithContext(ctx context.Context) MongodbrInsertOneOption {
	return func(mfoo *MongodbrInsertOneOptions) {
		mfoo.WithCtx = ctx
	}
}

// #endregion

// InsertOneOptions with context
type MongodbrInsertManyOptions struct {
	*options.InsertManyOptions
	WithContextOptions
}

type MongodbrInsertManyOption func(*MongodbrInsertManyOptions)

// #region MongodbrInsertManyOption members

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
func MongodbrInsertManyOptionWithContext(ctx context.Context) MongodbrInsertManyOption {
	return func(mfoo *MongodbrInsertManyOptions) {
		mfoo.WithCtx = ctx
	}
}

// #endregion

// DeleteOptions with context
type MongodbrDeleteOptions struct {
	*options.DeleteOptions
	WithContextOptions
}

type MongodbrDeleteOption func(*MongodbrDeleteOptions)

// #region MongodbrDeleteOption Members

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
func MongodbrDeleteOptionWithContext(ctx context.Context) MongodbrDeleteOption {
	return func(mfoo *MongodbrDeleteOptions) {
		mfoo.WithCtx = ctx
	}
}

// #endregion
