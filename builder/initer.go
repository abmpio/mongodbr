package builder

func init() {
	_opList[op_array_addToSet] = &Op{name: op_array_addToSet}
	_opList[op_array_pop] = &Op{name: op_array_pop}
	_opList[op_array_pull] = &Op{name: op_array_pull}
	_opList[op_array_push] = &Op{name: op_array_push}
	_opList[op_array_pullAll] = &Op{name: op_array_pullAll}
	_opList[op_array_elemMatch] = &Op{name: op_array_elemMatch}

	_opList[setKey] = &Op{name: setKey}

	_opList[op_comparison_eq] = &Op{name: op_comparison_eq}
	_opList[op_comparison_gt] = &Op{name: op_comparison_gt}
	_opList[op_comparison_gte] = &Op{name: op_comparison_gte}
	_opList[op_comparison_in] = &Op{name: op_comparison_in}
	_opList[op_comparison_lt] = &Op{name: op_comparison_lt}
	_opList[op_comparison_lte] = &Op{name: op_comparison_lte}
	_opList[op_comparison_ne] = &Op{name: op_comparison_ne}
	_opList[op_comparison_nin] = &Op{name: op_comparison_nin}
	_opList[op_comparison_regex] = &Op{name: op_comparison_regex}

	_opList[op_comparison_exists] = &Op{name: op_comparison_exists}
	_opList[op_comparison_type] = &Op{name: op_comparison_type}

	_opList[op_comparison_and] = &Op{name: op_comparison_and}
	_opList[op_comparison_not] = &Op{name: op_comparison_not}
	_opList[op_comparison_nor] = &Op{name: op_comparison_nor}
	_opList[op_comparison_or] = &Op{name: op_comparison_or}
}
