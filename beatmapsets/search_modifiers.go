package beatmapsets

type SearchOperator int
const (
	OpEquals SearchOperator = iota
	OpNotEquals
	OpGreaterThan
	OpGreaterThanOrEqual
	OpLessThan
	OpLessThanOrEqual

	// Only should be used for keywords
	OpNoop
)

func (q SubqueryParameter) Equals() SubqueryParameter {
	q.Operator = OpEquals
	return q
}

func (q SubqueryParameter) NotEquals() SubqueryParameter {
	q.Operator = OpNotEquals
	return q
}

func (q SubqueryParameter) GreaterThan() SubqueryParameter {
	q.Operator = OpGreaterThan
	return q
}

func (q SubqueryParameter) GreaterThanOrEqual() SubqueryParameter {
	q.Operator = OpGreaterThanOrEqual
	return q
}

func (q SubqueryParameter) LessThan() SubqueryParameter {
	q.Operator = OpLessThan
	return q
}

func (q SubqueryParameter) LessThanOrEqual() SubqueryParameter {
	q.Operator = OpLessThanOrEqual
	return q
}

func (q SubqueryParameter) LTE() SubqueryParameter { return q.LessThanOrEqual() }

func (q SubqueryParameter) GTE() SubqueryParameter { return q.GreaterThanOrEqual() }

func (op SearchOperator) String() string {
	switch (op) {
	case OpEquals:
		return "="
	case OpNotEquals:
		return "!="
	case OpGreaterThan:
		return ">"
	case OpLessThan:
		return "<"
	case OpGreaterThanOrEqual:
		return ">="
	case OpLessThanOrEqual:
		return "<="
	default:
		return ""
	}
}
