package two_pointers

import (
	"testing"
)

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Example 1:


Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:

Input: height = [4,2,0,3,2,5]
Output: 9


Constraints:

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/
// 0,1,0,2,1,0,1,3,2,1,2,1
// 0 1 2 3 4 5 6 7 8 9 10 11
func trap(height []int) int {
	answ := 0
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= leftMax {
				leftMax = height[l]
			} else {
				answ += leftMax - height[l]
			}
			l++
		} else {
			if height[r] >= rightMax {
				rightMax = height[r]
			} else {
				answ += rightMax - height[r]
			}
			r--
		}
	}

	return answ
}

func Test_trap(t *testing.T) {
	tcs := map[string]struct {
		input  []int
		expect int
	}{
		"1": {
			input:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expect: 6,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			answ := trap(tc.input)
			if answ != tc.expect {
				t.Errorf("got %d, want %d", answ, tc.expect)
			}
		})
	}
}
