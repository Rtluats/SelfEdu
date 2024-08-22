package stack

import "fmt"

type Stack struct {
	dataStore []int
}

func NewStack() *Stack {
	return &Stack{make([]int, 0)}
}

func (s *Stack) Push(data int) {
	s.dataStore = append(s.dataStore, data)
}

func (s *Stack) Pop() (int, error) {
	if len(s.dataStore) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	val := s.dataStore[len(s.dataStore)-1]
	s.dataStore = s.dataStore[:len(s.dataStore)-1]
	return val, nil
}
