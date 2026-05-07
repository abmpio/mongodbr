package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongodbrBulkWriteOptions struct {
	*options.BulkWriteOptions
	WithContextOptions
}

func (o *MongodbrBulkWriteOptions) List() []func(*options.BulkWriteOptions) error {
	if o.BulkWriteOptions == nil {
		o.BulkWriteOptions = &options.BulkWriteOptions{}
	}
	return asOptionLister(o.BulkWriteOptions).List()
}

type MongodbrBulkWriteOption func(*MongodbrBulkWriteOptions)

// merge MongodbrFindOption list and return one *MongodbrBulkWriteOptions
func MergeMongodbrBulkWriteOption(opts ...MongodbrBulkWriteOption) *MongodbrBulkWriteOptions {
	o := &MongodbrBulkWriteOptions{
		BulkWriteOptions: &options.BulkWriteOptions{},
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// MongodbrCountOptions with context
func MongodbrBulkWriteOptionWithContext(ctx context.Context) MongodbrBulkWriteOption {
	return func(mco *MongodbrBulkWriteOptions) {
		mco.WithCtx = ctx
	}
}
