package stack

import (
	"fmt"
	"testing"
)

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:

Input: n = 1
Output: ["()"]

Constraints:

	1 <= n <= 8
*/

func generateParenthesis(n int) []string {
	answ := make([]string, 0)
	//var backtrackWithStack func([]string)
	//backtrackWithStack = func(curr []string) {
	//	if len(curr) == (2 * n) {
	//		if isStillValid(curr, true) {
	//			answ = append(answ, strings.Join(curr, ""))
	//			return
	//		}
	//		return
	//	}
	//
	//	if !isStillValid(curr, false) {
	//		return
	//	}
	//
	//	for _, v := range []string{"(", ")"} {
	//		curr = append(curr, v)
	//		backtrackWithStack(curr)
	//		curr = curr[:len(curr)-1]
	//	}
	//}

	var backtraking func(string, int, int)
	backtraking = func(curr string, left, right int) {
		if len(curr) == 2*n {
			answ = append(answ, curr)
		}

		if left < n {
			backtraking(curr+"(", left+1, right)
		}
		if right < left {
			backtraking(curr+")", left, right+1)
		}
	}

	backtraking("", 0, 0)

	return answ
}

func isStillValid(s []string, final bool) bool {
	stack := make([]string, 0, len(s))
	for _, v := range s {
		if len(stack) == 0 || v == "(" {
			stack = append(stack, v)
			continue
		}

		lst := stack[len(stack)-1]

		if v == ")" && lst == "(" {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if final && len(stack) > 0 {
		return false
	}

	return true
}

func TestG(t *testing.T) {
	tcs := map[string]struct {
		actual   int
		expected []string
	}{
		"1": {
			actual:   1,
			expected: []string{"()"},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			a := generateParenthesis(tc.actual)
			fmt.Println(a)
		})
	}
}
