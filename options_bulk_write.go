package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbrBulkWriteOptions struct {
	*options.BulkWriteOptions
	WithContextOptions
}

type MongodbrBulkWriteOption func(*MongodbrBulkWriteOptions)

// merge MongodbrFindOption list and return one *MongodbrBulkWriteOptions
func MergeMongodbrBulkWriteOption(opts ...MongodbrBulkWriteOption) *MongodbrBulkWriteOptions {
	o := &MongodbrBulkWriteOptions{
		BulkWriteOptions: options.BulkWrite(),
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
