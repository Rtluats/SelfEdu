package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	tcs := map[string]struct {
		input    []int
		expected []int
	}{
		"push_one_elem": {
			input:    []int{1},
			expected: []int{1},
		},
		"push_many_elems": {
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			stack := NewStack()

			for _, v := range tc.input {
				stack.Push(v)
			}

			for _, e := range tc.expected {
				a, err := stack.Pop()
				require.NoError(t, err)
				require.Equal(t, e, a)
			}
			_, err := stack.Pop()
			require.Error(t, err)
			require.Equal(t, 0, len(stack.dataStore))
		})
	}
}

func TestStackPopWithZeroElement(t *testing.T) {
	stack := NewStack()
	_, err := stack.Pop()
	require.Equal(t, 0, len(stack.dataStore))
	require.Error(t, err)
}
