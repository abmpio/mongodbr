package builder

// build filter for regex
func Filter_Regex(filter map[string]interface{}, key string, v interface{}) {
	filter[key] = map[string]interface{}{"$regex": v, "$options": "i"}
}
