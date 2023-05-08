package builder

const (
	//https://www.mongodb.com/docs/manual/reference/operator/update-array/

	// Adds elements to an array only if they do not already exist in the set.
	op_array_addToSet string = "$addToSet"
	// Removes the first or last item of an array.
	op_array_pop string = "$pop"
	// Removes all array elements that match a specified query.
	op_array_pull string = "$pull"
	// Adds an item to an array.
	op_array_push string = "$push"
	// Removes all matching values from an array.
	op_array_pullAll string = "$pullAll"
)

func (l *OpList) AddToSet() string {
	return op_array_addToSet
}

func (l *OpList) Pop() string {
	return op_array_pop
}

func (l *OpList) Pull() string {
	return op_array_pull
}

func (l *OpList) Push() string {
	return op_array_push
}

func (l *OpList) PullAll() string {
	return op_array_pullAll
}
