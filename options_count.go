package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbrCountOptions struct {
	*options.CountOptions
	WithContextOptions
}

type MongodbrCountOption func(*MongodbrCountOptions)

// merge MongodbrCountOption list and return one *MongodbrCountOptions
func MergeMongodbrCountOption(opts ...MongodbrCountOption) *MongodbrCountOptions {
	o := &MongodbrCountOptions{
		CountOptions: options.Count(),
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// MongodbrCountOptions with context
func MongodbrCountOptionWithContext(ctx context.Context) MongodbrCountOption {
	return func(mco *MongodbrCountOptions) {
		mco.WithCtx = ctx
	}
}
