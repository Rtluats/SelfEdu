package stack

import (
	"strconv"
	"testing"
)

/*
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

    The valid operators are '+', '-', '*', and '/'.
    Each operand may be an integer or another expression.
    The division between two integers always truncates toward zero.
    There will not be any division by zero.
    The input represents a valid arithmetic expression in a reverse polish notation.
    The answer and all the intermediate calculations can be represented in a 32-bit integer.



Example 1:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Example 2:

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Example 3:

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22



Constraints:

    1 <= tokens.length <= 104
    tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].

*/

func evalRPN(tokens []string) int {
	vars := make([]int, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		// check is it int
		if _, err := strconv.Atoi(v); err == nil {
			v, _ := strconv.Atoi(v)
			vars = append(vars, v)
			continue
		}
		// this is sign
		prev1, prev2 := vars[len(vars)-1], vars[len(vars)-2]
		temp := 0
		vars = vars[:len(vars)-2]

		switch v {
		case "+":
			temp = prev2 + prev1
		case "-":
			temp = prev2 - prev1
		case "*":
			temp = prev2 * prev1
		case "/":
			temp = prev2 / prev1
		}
		vars = append(vars, temp)
	}

	return vars[0]
}

func TestRPN(t *testing.T) {
	tcs := map[string]struct {
		actual []string
		expect int
	}{
		"1": {
			actual: []string{"4", "13", "5", "/", "+"},
			expect: 6,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			act := evalRPN(tc.actual)
			if act != tc.expect {
				t.Errorf("got %d, want %d", act, tc.expect)
			}
		})
	}
}
