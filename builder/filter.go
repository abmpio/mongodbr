package builder

import "go.mongodb.org/mongo-driver/bson"

// build filter for regex
func Filter_Regex(filter map[string]interface{}, key string, v interface{}) map[string]interface{} {
	filter[key] = map[string]interface{}{"$regex": v, "$options": "i"}
	return filter
}

// build filter for regex and return this filter as bson.E
func Filter_RegexToBsonE(key string, v interface{}) bson.E {
	filter := bson.E{
		Key: key,
		Value: bson.M{
			"$regex": v, "$options": "i",
		},
	}
	return filter
}

// build filter for regex and return this filter as bson.M
func Filter_RegexToBsonM(key string, v interface{}) bson.M {
	filter := bson.M{
		key: bson.M{
			"$regex": v, "$options": "i",
		},
	}
	return filter
}

// 将一个key增加到or条件列表中
func Filter_Or(filter []map[string]interface{}, key string, v interface{}) []map[string]interface{} {
	filter = append(filter, bson.M{
		key: v,
	})
	return filter
}
