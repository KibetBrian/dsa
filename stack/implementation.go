package stack

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

func (s *Stack) Peek() int{
	return s.values[len(s.values)-1]
}