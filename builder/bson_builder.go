package builder

import "go.mongodb.org/mongo-driver/bson"

const (
	setKey   string = "$set"
	unsetKey string = "$unset"
)

type BsonBuilder struct {
	bson bson.M
}

func NewBsonBuilder() *BsonBuilder {
	return &BsonBuilder{}
}

// 增加$set类型的值
func (b *BsonBuilder) NewOrUpdateSet(v interface{}) *BsonBuilder {
	b.ensureBson()
	b.bson[setKey] = v
	return b
}

// 构建一个$set类型的值，用于设置字段值
func NewOrUpdateSetBsonBuilder(v interface{}) *BsonBuilder {
	b := NewBsonBuilder()
	b.NewOrUpdateSet(v)
	return b
}

// 增加$unset类型的值,用于移除字段值
func UnsetBsonBuilder(fields []string) *BsonBuilder {
	b := NewBsonBuilder().ensureBson()
	unsetValue := bson.M{}
	b.bson[unsetKey] = unsetValue
	// iterate for fields
	for _, eachField := range fields {
		unsetValue[eachField] = ""
	}
	return b
}

func (b *BsonBuilder) ensureBson() *BsonBuilder {
	if b.bson == nil {
		b.bson = bson.M{}
	}
	return b
}

func (b *BsonBuilder) ToValue() bson.M {
	return b.bson
}
