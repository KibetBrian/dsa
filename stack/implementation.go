package stack

import "errors"

//Implementing stack using slice
type Stack struct {
	values []int
}

func (s *Stack) Push(i int) {
	s.values = append(s.values, i)
}

func (s *Stack) Pop() int {
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val
}

func (s *Stack) Peek() (int, error){
	if len(s.values) == 0{
		return -1, errors.New("Stack is empty");
	}
	return s.values[len(s.values)-1], nil
}

func (s *Stack) Size() int{
	return len(s.values)
}