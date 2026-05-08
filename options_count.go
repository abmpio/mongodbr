package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongodbrCountOptions struct {
	*options.CountOptions
	WithContextOptions
}

func (o *MongodbrCountOptions) List() []func(*options.CountOptions) error {
	if o.CountOptions == nil {
		o.CountOptions = &options.CountOptions{}
	}
	return asOptionLister(o.CountOptions).List()
}

type MongodbrCountOption func(*MongodbrCountOptions)

// merge MongodbrCountOption list and return one *MongodbrCountOptions
func MergeMongodbrCountOption(opts ...MongodbrCountOption) *MongodbrCountOptions {
	o := &MongodbrCountOptions{
		CountOptions: &options.CountOptions{},
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
