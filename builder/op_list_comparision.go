package builder

const (
	//https://www.mongodb.com/docs/manual/reference/operator/query-comparison/

	//Matches values that are equal to a specified value.
	op_comparison_eq string = "$eq"
	//Matches values that are greater than a specified value.
	op_comparison_gt string = "$gt"
	//Matches values that are greater than or equal to a specified value.
	op_comparison_gte string = "$gte"
	//Matches any of the values specified in an array.
	op_comparison_in string = "$in"
	//Matches values that are less than a specified value.
	op_comparison_lt string = "$lt"
	//Matches values that are less than or equal to a specified value.
	op_comparison_lte string = "$lte"
	//Matches all values that are not equal to a specified value.
	op_comparison_ne string = "$ne"
	//Matches none of the values specified in an array.
	op_comparison_nin string = "$nin"

	op_comparison_regex string = "$regex"
)

type Op struct {
	name string
}

func (op *Op) String() string {
	return op.name
}

var (
	_opList = map[string]*Op{}
)

func Op_Eq() *Op {
	return _opList[op_comparison_eq]
}

func Op_Gt() *Op {
	return _opList[op_comparison_gt]
}

func Op_Gte() *Op {
	return _opList[op_comparison_gte]
}

func Op_In() *Op {
	return _opList[op_comparison_in]
}

func Op_Lt() *Op {
	return _opList[op_comparison_lt]
}

func Op_Lte() *Op {
	return _opList[op_comparison_lte]
}

func Op_Ne() *Op {
	return _opList[op_comparison_ne]
}

func Op_Nin() *Op {
	return _opList[op_comparison_nin]
}

func Op_Regex() *Op {
	return _opList[op_comparison_regex]
}
