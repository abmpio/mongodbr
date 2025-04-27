package mongodbr

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DefaultConfiguration                                   = NewConfiguration()
	_cachedClientOptions map[string]*options.ClientOptions = make(map[string]*options.ClientOptions)
	// 是否忽略uuid的自定义解码器
	_ignoreUUIDDecoder = true
	// 是否忽略time.Time的自定义解码器
	_ignoreTimeDecoder = true
)

// enable mongodb monitor
func EnableMongodbMonitor() func(*options.ClientOptions) {
	return func(co *options.ClientOptions) {
		monitor := &event.CommandMonitor{
			Started: func(_ context.Context, e *event.CommandStartedEvent) {
				log.Println(e.Command.String())
			},
			Succeeded: func(ctx context.Context, e *event.CommandSucceededEvent) {
				log.Println(e.Reply.String())
			},
			Failed: func(ctx context.Context, e *event.CommandFailedEvent) {
				log.Println("mongodb error:", e.Failure)
			},
		}

		co.SetMonitor(monitor)
	}
}

// 在创建Client时是否忽略uuid的自定义解码器
func IgnoreUUIDDecoder(ignore bool) {
	_ignoreTimeDecoder = ignore
}

// 在创建Client时是否忽略time.Time的自定义解码器,(primitive.DateTime -> time.Time)
func IgnoreTimeDecoder(ignore bool) {
	_ignoreTimeDecoder = ignore
}

func DefaultClientOptions() *options.ClientOptions {
	return _cachedClientOptions[DefaultAlias]
}

// get client options by key
func GetClientOptions(key string) *options.ClientOptions {
	clientOptions, ok := _cachedClientOptions[key]
	if !ok {
		return nil
	}
	return clientOptions
}

type Configuration struct {
	QueryTimeout time.Duration

	//创建一条新的记录,并返回这条记录的指针地址
	createItemFunc func() interface{}
	//查询时设置默认的排序
	setDefaultSort func(*options.FindOptions) *options.FindOptions
}

func CreateContextAndCancel(c *Configuration) (context.Context, context.CancelFunc) {
	if c == nil || c.QueryTimeout <= 0 {
		ctx := context.TODO()
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(context.Background(), c.QueryTimeout)
}

// 使用parent的context来创建一个context
func CreateContextAndCancelWith(c *Configuration, ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		return CreateContextAndCancel(c)
	}
	// 包含了ctx
	if c == nil || c.QueryTimeout <= 0 {
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(ctx, c.QueryTimeout)
}

// // 使用parent的context来创建一个context
// func CreateContextWith(c *Configuration, ctx context.Context) (context.Context, context.CancelFunc) {
// 	if ctx == nil {
// 		return CreateContext(c)
// 	}
// 	// 包含了ctx
// 	if c == nil || c.QueryTimeout <= 0 {
// 		return context.WithCancel(ctx)
// 	}
// 	return context.WithTimeout(ctx, c.QueryTimeout)
// }

func (c *Configuration) safeCreateItem() interface{} {
	if c.createItemFunc == nil {
		return make(map[string]interface{})
	}
	return c.createItemFunc()
}

func NewConfiguration() *Configuration {
	return &Configuration{
		QueryTimeout: 120 * time.Second,
	}
}

type RepositoryOption func(*Configuration)

func WithDefaultSort(defaultSortFunc func(*options.FindOptions) *options.FindOptions) RepositoryOption {
	return func(configuration *Configuration) {
		configuration.setDefaultSort = defaultSortFunc
	}
}

func WithCreateItemFunc(createItemFunc func() interface{}) RepositoryOption {
	return func(configuration *Configuration) {
		configuration.createItemFunc = createItemFunc
	}
}
