package builder

import (
	"github.com/abmpio/libx/lang"
)

type BsonBuilderOption func(*BsonBuilder)

// new mongodbBuilder.BsonBuilder with option
func NewBsonBuilderWithOption(opts ...BsonBuilderOption) *BsonBuilder {
	bsonBuilder := NewBsonBuilder()
	for _, eachOpt := range opts {
		eachOpt(bsonBuilder)
	}
	return bsonBuilder
}

// set lastModificationTime field value to now
func BsonBuilderOptionWithLastModificationTime() BsonBuilderOption {
	return func(bb *BsonBuilder) {
		bb.AppendSetField("lastModificationTime", lang.NowToPtr())
	}
}

// set lastModifierId field value to now
func BsonBuilderOptionWithLastModifierId(userId string) BsonBuilderOption {
	return func(bb *BsonBuilder) {
		bb.AppendSetField("lastModifierId", userId)
	}
}

// set lastModifierId field value to now
func BsonBuilderOptionWith(fieldName string, v interface{}) BsonBuilderOption {
	return func(bb *BsonBuilder) {
		bb.AppendSetField(fieldName, v)
	}
}
