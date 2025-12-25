package generic2

import "errors"

type Stack[T any] struct {
	Items []T
}

// Push methode de la structure
func (s *Stack[T]) Push(v T) {
	s.Items = append(s.Items, v)
}

// Pop methode de la structure
func (s *Stack[T]) Pop() (T, error) {
	if len(s.Items) == 0 {
		var negative T
		return negative, errors.New("stack is empty")
	}
	var lastItem T = s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return lastItem, nil
}

type Structure[T any] struct {
	VarName1 T
	VarName2 T
	VarName3 T
}

// Operation method of structure
func (s *Structure[T]) Operation(v1 T, v2 T, v3 T) {
	s.VarName1 = v1
	s.VarName2 = v2
	s.VarName3 = v3
}
