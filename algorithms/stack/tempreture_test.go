package stack

import (
	"reflect"
	"testing"
)

/*
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have
to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.



Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]



Constraints:

    1 <= temperatures.length <= 105
    30 <= temperatures[i] <= 100

*/
// {73, 74, 75, 71, 69, 72, 76, 73}
func dailyTemperatures(temperatures []int) []int {
	answ := make([]int, len(temperatures))
	stack := make([][2]int, 0, len(temperatures))
	// 				  {73, 74, 75, 71, 69, 72, 76, 73}
	// 				  stack
	for i, v := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, [2]int{i, v})
			continue
		}

		top := stack[len(stack)-1]
		for top[1] < v {
			answ[top[0]] = i - top[0]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			top = stack[len(stack)-1]
		}

		stack = append(stack, [2]int{i, v})
	}

	return answ
}

func TestTemp(t *testing.T) {
	tcs := map[string]struct {
		temperatures []int
		expected     []int
	}{
		"1": {
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			answ := dailyTemperatures(tc.temperatures)
			if !reflect.DeepEqual(answ, tc.expected) {
				t.Errorf("got %v, want %v", answ, tc.expected)
			}
		})
	}
}
