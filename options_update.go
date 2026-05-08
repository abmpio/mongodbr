package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// FindOneAndUpdateOptions with context
type MongodbrFindOneAndUpdateOptions struct {
	*options.FindOneAndUpdateOptions
	WithContextOptions
}

func (o *MongodbrFindOneAndUpdateOptions) List() []func(*options.FindOneAndUpdateOptions) error {
	if o.FindOneAndUpdateOptions == nil {
		o.FindOneAndUpdateOptions = &options.FindOneAndUpdateOptions{}
	}
	return asOptionLister(o.FindOneAndUpdateOptions).List()
}

type MongodbrFindOneAndUpdateOption func(*MongodbrFindOneAndUpdateOptions)

// #region MongodbrFindOneAndUpdateOption members

// merge MongodbrFindOneAndUpdateOption list and return one *MongodbrFindOneAndUpdateOptions
func MergeMongodbrFindOneAndUpdateOption(opts ...MongodbrFindOneAndUpdateOption) *MongodbrFindOneAndUpdateOptions {
	o := &MongodbrFindOneAndUpdateOptions{
		FindOneAndUpdateOptions: &options.FindOneAndUpdateOptions{},
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
	*options.UpdateOneOptions
	*options.UpdateManyOptions
	WithContextOptions
}

type MongodbrUpdateOption func(*MongodbrUpdateOptions)

// #region MongodbrUpdateOption members

// merge MongodbrUpdateOption list and return one *MongodbrUpdateOption
func MergeMongodbrUpdateOption(opts ...MongodbrUpdateOption) *MongodbrUpdateOptions {
	o := &MongodbrUpdateOptions{
		UpdateOneOptions:  &options.UpdateOneOptions{},
		UpdateManyOptions: &options.UpdateManyOptions{},
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

func (o *MongodbrReplaceOptions) List() []func(*options.ReplaceOptions) error {
	if o.ReplaceOptions == nil {
		o.ReplaceOptions = &options.ReplaceOptions{}
	}
	return asOptionLister(o.ReplaceOptions).List()
}

type MongodbrReplaceOption func(*MongodbrReplaceOptions)

// #region MongodbrReplaceOption Members

// merge MongodbrReplaceOption list and return one *MongodbrReplaceOptions
func MergeMongodbrReplaceOption(opts ...MongodbrReplaceOption) *MongodbrReplaceOptions {
	o := &MongodbrReplaceOptions{
		ReplaceOptions: &options.ReplaceOptions{},
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

func (o *MongodbrInsertOneOptions) List() []func(*options.InsertOneOptions) error {
	if o.InsertOneOptions == nil {
		o.InsertOneOptions = &options.InsertOneOptions{}
	}
	return asOptionLister(o.InsertOneOptions).List()
}

type MongodbrInsertOneOption func(*MongodbrInsertOneOptions)

// #region MongodbrInsertOneOption members

// merge MongodbrInsertOneOption list and return one *MongodbrInsertOneOptions
func MergeMongodbrInsertOneOption(opts ...MongodbrInsertOneOption) *MongodbrInsertOneOptions {
	o := &MongodbrInsertOneOptions{
		InsertOneOptions: &options.InsertOneOptions{},
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

func (o *MongodbrInsertManyOptions) List() []func(*options.InsertManyOptions) error {
	if o.InsertManyOptions == nil {
		o.InsertManyOptions = &options.InsertManyOptions{}
	}
	return asOptionLister(o.InsertManyOptions).List()
}

type MongodbrInsertManyOption func(*MongodbrInsertManyOptions)

// #region MongodbrInsertManyOption members

// merge MongodbrInsertManyOption list and return one *MongodbrInsertManyOptions
func MergeMongodbrInsertManyOption(opts ...MongodbrInsertManyOption) *MongodbrInsertManyOptions {
	o := &MongodbrInsertManyOptions{
		InsertManyOptions: &options.InsertManyOptions{},
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
	*options.DeleteOneOptions
	*options.DeleteManyOptions
	WithContextOptions
}

type MongodbrDeleteOption func(*MongodbrDeleteOptions)

// #region MongodbrDeleteOption Members

// merge MongodbrDeleteOption list and return one *MongodbrDeleteOptions
func MergeMongodbrDeleteOption(opts ...MongodbrDeleteOption) *MongodbrDeleteOptions {
	o := &MongodbrDeleteOptions{
		DeleteOneOptions:  &options.DeleteOneOptions{},
		DeleteManyOptions: &options.DeleteManyOptions{},
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
