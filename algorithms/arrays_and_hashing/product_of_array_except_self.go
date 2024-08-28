package arrays_and_hashing

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:

1) 2 <= nums.length <= 105
2) -30 <= nums[i] <= 30
3) The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer

*/

// 0(n^2)
func bruteforce(nums []int) []int {
	answ := []int{}
	for i, _ := range nums {
		val := 1
		for j, _ := range nums {
			if i != j {
				val *= nums[j]
			}
		}
		answ = append(answ, val)
	}
	return answ
}

// O(3n)== O(n), O(n)
func On(nums []int) []int {
	answ := make([]int, len(nums))
	prefixes := make([]int, len(nums))
	suffixes := make([]int, len(nums))

	for i, _ := range prefixes {
		if i == 0 {
			prefixes[i] = 1
			continue
		}

		prefixes[i] = prefixes[i-1] * nums[i-1]
	}

	for i := len(suffixes) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			suffixes[i] = 1
			continue
		}
		suffixes[i] = suffixes[i+1] * nums[i+1]
	}

	for i, _ := range answ {
		answ[i] = prefixes[i] * suffixes[i]
	}

	return answ
}

// O(n), O(1)
func O1n(nums []int) []int {
	answ := make([]int, len(nums))
	// prefix
	for i, _ := range answ {
		if i == 0 {
			answ[i] = 1
			continue
		}

		answ[i] = answ[i-1] * nums[i-1]
	}

	//suffix
	k := 1 // actual suffix
	for i := len(answ) - 2; i >= 0; i-- {
		k *= nums[i+1]
		answ[i] *= k
	}

	return answ
}
