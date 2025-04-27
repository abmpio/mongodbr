package mongodbr

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOneAndUpdateOptions with context
type MongodbrFindOneAndUpdateOptions struct {
	*options.FindOneAndUpdateOptions
	WithContextOptions
}

type MongodbrFindOneAndUpdateOption func(*MongodbrFindOneAndUpdateOptions)

// UpdateOptions with context
type MongodbrUpdateOptions struct {
	*options.UpdateOptions
	WithContextOptions
}

type MongodbrUpdateOption func(*MongodbrUpdateOptions)

// ReplaceOptions with context
type MongodbrReplaceOptions struct {
	*options.ReplaceOptions
	WithContextOptions
}

type MongodbrReplaceOption func(*MongodbrReplaceOptions)

// InsertOneOptions with context
type MongodbrInsertOneOptions struct {
	*options.InsertOneOptions
	WithContextOptions
}

type MongodbrInsertOneOption func(*MongodbrInsertOneOptions)

// InsertOneOptions with context
type MongodbrInsertManyOptions struct {
	*options.InsertManyOptions
	WithContextOptions
}

type MongodbrInsertManyOption func(*MongodbrInsertManyOptions)

// DeleteOptions with context
type MongodbrDeleteOptions struct {
	*options.DeleteOptions
	WithContextOptions
}

type MongodbrDeleteOption func(*MongodbrDeleteOptions)
