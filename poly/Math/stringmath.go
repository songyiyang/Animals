package Math

type StringMath struct {
}

var _ Math = (*StringMath)(nil)

func (s *StringMath) Add(a interface{}, b interface{}) interface{} {
	aStr := a.(string)
	bStr := b.(string)

	return aStr + bStr
}
