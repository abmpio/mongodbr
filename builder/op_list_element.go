package builder

import "go.mongodb.org/mongo-driver/bson"

const (
	//https://www.mongodb.com/docs/manual/reference/operator/query-element/

	//Matches documents that have the specified field.
	op_comparison_exists string = "$exists"
	//Selects documents if a field is of the specified type.
	op_comparison_type string = "$type"
)

func Op_Exists() *Op {
	return _opList[op_comparison_exists]
}

func Op_Type() *Op {
	return _opList[op_comparison_type]
}

// append $exists filter expression to bson.D
func AppendExistsFilterToBsonD(fieldName string, existsValue bool, d bson.D) bson.D {
	v := bson.E{Key: fieldName, Value: bson.D{
		{Key: "$exists", Value: existsValue},
	}}
	d = append(d, v)
	return d
}
