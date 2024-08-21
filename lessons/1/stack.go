package stack

type Stack struct {
	dataStore []interface{}
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0)}
}

func (s *Stack) Push(data interface{}) {
	s.dataStore = append(s.dataStore, data)
}

func (s *Stack) Pop() interface{} {
	if len(s.dataStore) == 0 {
		return nil
	}
	val := s.dataStore[len(s.dataStore)-1]
	s.dataStore = s.dataStore[:len(s.dataStore)-1]
	return val
}
