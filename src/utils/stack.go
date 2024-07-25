package utils

type Stack []interface{}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) pop() interface{} {
	if s.isEmpty() {
		return nil
	}
	top := len(*s) - 1
	value := (*s)[top]
	*s = (*s)[:top]
	return value
}
