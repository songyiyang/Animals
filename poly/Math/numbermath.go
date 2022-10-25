package Math

type NumberMath struct {
}

var _ Math = (*NumberMath)(nil)

func (s *NumberMath) Add(a interface{}, b interface{}) interface{} {
	aInt := a.(int)
	bInt := b.(int)

	return aInt + bInt
}
