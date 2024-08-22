package stack

import "fmt"

type Stack []int

func NewStack() Stack {
	return make(Stack, 0)
}

func (s *Stack) Push(data int) {
	*s = append(*s, data)
}

func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, nil
}
