package arrays_and_hashing

import (
	"sort"
)

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109

*/

// O(nlogn), O(n)
func nlogn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	set := make(map[int]bool, len(nums))

	for _, n := range nums {
		set[n] = false
	}

	if len(set) == 1 {
		return 1
	}

	keys := make([]int, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	curr, m := 1, 0
	for i := 1; i < len(keys); i++ {
		if keys[i]-keys[i-1] == 1 {
			curr += 1
			continue
		}
		if curr > m {
			m = curr
		}
		curr = 1
	}

	if curr > m {
		m = curr
	}

	return m
}

// o(n)
func n(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, n := range nums {
		set[n] = true
	}
	m := 0
	for v := range set {
		if _, ok := set[v-1]; !ok { // begin of sequence
			currNum := v
			curr := 1
			// check if next exist
			for exist := set[currNum+1]; exist; exist = set[currNum+1] { // step in sequence
				currNum += 1 // if exist let's add this, and then check if exist next
				curr += 1
			}
			if curr > m {
				m = curr
			}
		}
	}

	return m
}
