package mongodbr

import "go.mongodb.org/mongo-driver/mongo/options"

type MongodbrAggregateOptions struct {
	*options.AggregateOptions
	WithContextOptions
}

// AggregateOptions handler pipeline
type MongodbrAggregateOption func(*MongodbrAggregateOptions)
