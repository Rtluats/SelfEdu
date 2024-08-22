package stack

type Stack struct {
	dataStore []any
}

func NewStack() *Stack {
	return &Stack{make([]any, 0)}
}

func (s *Stack) Push(data any) {
	s.dataStore = append(s.dataStore, data)
}

func (s *Stack) Pop() any {
	if len(s.dataStore) == 0 {
		return nil
	}
	val := s.dataStore[len(s.dataStore)-1]
	s.dataStore = s.dataStore[:len(s.dataStore)-1]
	return val
}
