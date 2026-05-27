package stackdriver

type FilterOption func(filter *filter)

type filter struct {
	operator string

	leftOperand  string
	rightOperand string
}

func Eq(leftOperand string, rightOperand string) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}

func Neq(leftOperand string, rightOperand string) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}

func Matches(leftOperand string, rightOperand string) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}

func NotMatches(leftOperand string, rightOperand string) FilterOption {
	_ = "STUB: not implemented"
	return *new(FilterOption)
}
