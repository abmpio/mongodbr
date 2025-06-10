package builder

import (
	"github.com/abmpio/libx/lang/tuple"
	"go.mongodb.org/mongo-driver/bson"
)

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

// get $set value
// if not exist,then add $set value
func (b *BsonBuilder) SetOrDefault() bson.M {
	b.ensureBson()
	setValue, ok := b.bson[setKey].(bson.M)
	if ok && setValue != nil {
		return setValue
	}
	setValue = bson.M{}
	b.bson[setKey] = setValue
	return setValue
}

// get $unset value
// if not exist,then add $unset value
func (b *BsonBuilder) UnsetOrDefault() bson.M {
	b.ensureBson()
	unsetValue, ok := (b.bson[unsetKey]).(bson.M)
	if ok && unsetValue != nil {
		return unsetValue
	}
	unsetValue = bson.M{}
	b.bson[unsetKey] = unsetValue
	return unsetValue
}

// append field value to $set field
// V1:key V2:value
func (b *BsonBuilder) AppendSetField(fieldValueList ...tuple.T2[string, interface{}]) *BsonBuilder {
	if len(fieldValueList) <= 0 {
		return b
	}
	setValue := b.SetOrDefault()
	for _, eachValue := range fieldValueList {
		setValue[eachValue.V1] = eachValue.V2
	}
	return b
}

// append field to $unset value
func (b *BsonBuilder) AppendUnsetField(fieldList []string) *BsonBuilder {
	if len(fieldList) <= 0 {
		return b
	}
	unsetValue := b.UnsetOrDefault()
	for _, eachField := range fieldList {
		unsetValue[eachField] = ""
	}
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
