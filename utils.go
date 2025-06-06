package mongodbr

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConvertPrimitiveDToMapOptions struct {
	ConvertPrimitiveAToMap bool
	ConvertPrimitiveDToMap bool
}

func _bsonDToMap(d primitive.D) map[string]interface{} {
	m := make(map[string]interface{})
	for _, elem := range d {
		m[elem.Key] = elem.Value
	}
	return m
}

// bsonDToMap 将 bson.D 转换为 map[string]interface{}
func bsonDToMap(d primitive.D, opts ...ConvertPrimitiveDToMapOptions) map[string]interface{} {
	options := ConvertPrimitiveDToMapOptions{}
	if len(opts) > 0 {
		options = opts[len(opts)-1]
	}
	m := make(map[string]interface{})
	for _, elem := range d {
		elemValue := elem.Value
		if options.ConvertPrimitiveAToMap {
			array, ok := elem.Value.(primitive.A)
			if ok {
				// 将其转换为map[string]interface{}
				elemValue = ConvertPrimitiveAToMap(array)
			}
		}
		if options.ConvertPrimitiveDToMap {
			dValue, ok := elem.Value.(primitive.D)
			if ok {
				elemValue = _bsonDToMap(dValue)
			}
		}
		m[elem.Key] = elemValue
	}
	return m
}

// bsonAToMap 将 bson.A 转换为 map[string]interface{}
func bsonAToMap(a primitive.A) map[string]interface{} {
	result := make(map[string]interface{})
	// 遍历数组
	for _, item := range a {
		switch val := item.(type) {
		case bson.D:
			// bson.D -> map[string]interface{}
			m := bsonDToMap(val)
			for k, v := range m {
				result[k] = v
			}
		case bson.M:
			// bson.M 直接合并
			for k, v := range val {
				result[k] = v
			}
		case map[string]interface{}:
			// map[string]interface{} 直接合并
			for k, v := range val {
				result[k] = v
			}
		}
	}
	return result
}

func ConvertPrimitiveDToMap(d primitive.D, opts ...ConvertPrimitiveDToMapOptions) map[string]interface{} {
	return bsonDToMap(d, opts...)
}

// 将[]interface{}中的值转换为map[string]interface{}
func ConvertPrimitiveAToMap(v interface{}) map[string]interface{} {
	// 确保输入是 []interface{}
	a, ok := v.(primitive.A)
	if !ok {
		return make(map[string]interface{})
	}
	return bsonAToMap(a)
}

// convert value to map
// only when v is primitive.D or primitive.A
func ConvertValueToMap(v interface{}) map[string]interface{} {
	if v == nil {
		return nil
	}
	switch value := v.(type) {
	case primitive.D:
		return bsonDToMap(value)
	case primitive.A:
		return bsonAToMap(value)
	default:
		// other return nil
		return nil
	}
}
