package mongodbr

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertPrimitiveDToMap(doc primitive.D) map[string]interface{} {
	result := make(map[string]interface{})
	for _, elem := range doc {
		result[elem.Key] = elem.Value
	}
	return result
}
