package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongodbrAggregateOptions struct {
	*options.AggregateOptions
	WithContextOptions
}

func (o *MongodbrAggregateOptions) List() []func(*options.AggregateOptions) error {
	if o.AggregateOptions == nil {
		o.AggregateOptions = &options.AggregateOptions{}
	}
	return asOptionLister(o.AggregateOptions).List()
}

// AggregateOptions handler pipeline
type MongodbrAggregateOption func(*MongodbrAggregateOptions)

// merge MongodbrAggregateOption list and return one *MongodbrAggregateOptions
func MergeMongodbrAggregateOption(opts ...MongodbrAggregateOption) *MongodbrAggregateOptions {
	o := &MongodbrAggregateOptions{
		AggregateOptions: &options.AggregateOptions{},
	}
	for _, eachOpt := range opts {
		eachOpt(o)
	}
	return o
}

// MongodbrAggregateOption with context
func MongodbrAggregateOptionWithContext(ctx context.Context) MongodbrAggregateOption {
	return func(mco *MongodbrAggregateOptions) {
		mco.WithCtx = ctx
	}
}
