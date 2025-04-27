package mongodbr

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DefaultAlias                           = "default"
	_cachedClient map[string]*mongo.Client = make(map[string]*mongo.Client)
)

func DefaultClient() *mongo.Client {
	return _cachedClient[DefaultAlias]
}

// 构建默认的client
func SetupDefaultClient(uri string, opts ...func(*options.ClientOptions)) (*mongo.Client, error) {
	return RegistClient(DefaultAlias, uri, opts...)
}

func RegistClient(key string, uri string, opts ...func(*options.ClientOptions)) (*mongo.Client, error) {
	client, clientOptions, err := CreateClient(uri, opts...)
	if err != nil {
		return nil, err
	}
	_cachedClient[key] = client
	_cachedClientOptions[key] = options.MergeClientOptions(clientOptions)
	return client, nil
}

// get client by key
func GetClient(key string) *mongo.Client {
	client, ok := _cachedClient[key]
	if !ok {
		return nil
	}
	return client
}

func CreateClient(uri string, opts ...func(*options.ClientOptions)) (*mongo.Client, *options.ClientOptions, error) {
	mongoRegistry := bson.NewRegistry()
	continRegistry := false
	if !_ignoreUUIDDecoder {
		mongoRegistry.RegisterTypeEncoder(_tUUID, bsoncodec.ValueEncoderFunc(uuidEncodeValue))
		mongoRegistry.RegisterTypeDecoder(_tUUID, bsoncodec.ValueDecoderFunc(uuidDecodeValue))
		continRegistry = true
	}
	if !_ignoreTimeDecoder {
		// registryBuilder = registryBuilder.
		// 	RegisterTypeDecoder(reflect.TypeOf(time.Time{}), bsoncodec.ValueDecoderFunc(timeDecodeValue))
		// continRegistry = true
	}

	//测试能否连接
	clientOptions := options.Client().ApplyURI(uri)
	if continRegistry && mongoRegistry != nil {
		clientOptions.SetRegistry(mongoRegistry)
	}
	for _, eachOpt := range opts {
		eachOpt(clientOptions)
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("无法初始化mongodb,在连接到mongodb时出现异常,异常信息:%s", err.Error())
	}
	return client, clientOptions, nil
}
