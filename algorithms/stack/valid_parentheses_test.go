package stack

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true



Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	stack := make([]string, 0, len(s))
	openClose := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, string(ch))
		} else {
			if len(stack) == 0 {
				return false
			}
			lst := stack[len(stack)-1]
			if openClose[lst] != string(ch) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
