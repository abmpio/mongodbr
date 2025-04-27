package mongodbr

import "context"

type IWithContext interface {
	GetContext() context.Context
}

// 带context.Context实例的Options，如果不为空，将由函数上下文来处理
type WithContextOptions struct {
	WithCtx context.Context
}

var _ IWithContext = (*WithContextOptions)(nil)

// new WithContextOptions instance
func NewWithContextOptions() *WithContextOptions {
	return &WithContextOptions{}
}

func (o *WithContextOptions) MergeWithContextOptions(opts ...*WithContextOptions) *WithContextOptions {
	for _, eachOpt := range opts {
		o.WithCtx = eachOpt.WithCtx
	}
	return o
}

// #region IWithContext Members

func (o WithContextOptions) GetContext() context.Context {
	return o.WithCtx
}

// #endregion
