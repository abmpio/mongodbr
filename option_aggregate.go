package mongodbr

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbrAggregateOptions struct {
	*options.AggregateOptions
	WithContextOptions
}

// AggregateOptions handler pipeline
type MongodbrAggregateOption func(*MongodbrAggregateOptions)

// merge MongodbrAggregateOption list and return one *MongodbrAggregateOptions
func MergeMongodbrAggregateOption(opts ...MongodbrAggregateOption) *MongodbrAggregateOptions {
	o := &MongodbrAggregateOptions{
		AggregateOptions: options.Aggregate(),
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
