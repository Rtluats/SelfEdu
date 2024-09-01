package two_pointers

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/
func isalnum(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z')
}

func tolower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func ispolidroment(s string) bool {
	l, r := 0, len(s)-1
	for l < r {

		for len(s)-1 > l && !isalnum(s[l]) {
			l += 1
		}

		for r > 0 && !isalnum(s[r]) {
			r -= 1
		}

		if l > r {
			return true
		}

		if tolower(s[l]) != tolower(s[r]) {
			return false
		}
		l += 1
		r -= 1

	}

	return true
}
