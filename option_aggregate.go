package mongodbr

import (
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
