package two_pointers

import "sort"

/*

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.


Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func threeSum(nums []int) [][]int {
	res := make(map[[3]int]int)
	dups := make(map[int]bool)
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := dups[nums[i]]; !ok {
			dups[nums[i]] = true
			for j := i + 1; j < len(nums); j++ {
				component := -nums[i] - nums[j]
				if v, ok := seen[component]; ok && v == i {
					triplet := []int{nums[i], nums[j], component}
					sort.Ints(triplet)
					res[[3]int{triplet[0], triplet[1], triplet[2]}] = 1
				}
				seen[nums[j]] = i
			}
		}
	}
	keys := make([][]int, 0)
	for k := range res {
		keys = append(keys, []int{k[0], k[1], k[2]})
	}
	return keys
}

func threeSum2(nums []int) [][]int {

	res := make(map[[3]int]bool, len(nums))
	values := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		values[nums[i]] += 1
	}

	if c, _ := values[0]; c >= 3 && len(values) == 1 {
		return [][]int{
			{
				0, 0, 0,
			},
		}
	}

	if len(values) <= 1 {
		return [][]int{}
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-1 && nums[i] <= 0; i++ {
		for j := i + 1; j < len(nums); j++ {
			component := -nums[j] - nums[i]
			if _, ok := values[component]; ok {
				values[component] -= 1
				values[nums[i]] -= 1
				values[nums[j]] -= 1
				if values[component] >= 0 && values[nums[i]] >= 0 && values[nums[j]] >= 0 {
					triplet := []int{nums[i], nums[j], component}
					sort.Ints(triplet)
					res[[3]int{triplet[0], triplet[1], triplet[2]}] = true
				}
				values[component] += 1
				values[nums[i]] += 1
				values[nums[j]] += 1
			}
		}
	}

	keys := make([][]int, 0, len(res))
	for k := range res {
		keys = append(keys, []int{k[0], k[1], k[2]})
	}

	return keys
}
