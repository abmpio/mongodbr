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

// MongodbrCountOptions with context
func MongodbrCountOptionWithContext(ctx context.Context) MongodbrCountOption {
	return func(mco *MongodbrCountOptions) {
		mco.WithCtx = ctx
	}
}
