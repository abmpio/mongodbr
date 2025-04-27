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

// MongodbrCountOptions with context
func MongodbrBulkWriteOptionWithContext(ctx context.Context) MongodbrBulkWriteOption {
	return func(mco *MongodbrBulkWriteOptions) {
		mco.WithCtx = ctx
	}
}
