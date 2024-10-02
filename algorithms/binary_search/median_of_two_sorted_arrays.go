package binary_search

import "math"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.


Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	index1, index2 := 0, 0
//	merged := make([]int, 0, len(nums1)+len(nums2))
//
//	for index1 < len(nums1) && index2 < len(nums2) {
//		if nums1[index1] < nums2[index2] {
//			merged = append(merged, nums1[index1])
//			index1++
//		} else {
//			merged = append(merged, nums2[index2])
//			index2++
//		}
//	}
//
//	for index1 < len(nums1) {
//		merged = append(merged, nums1[index1])
//		index1++
//	}
//
//	for index2 < len(nums2) {
//		merged = append(merged, nums2[index2])
//		index2++
//	}
//
//	mid := len(merged) / 2
//	if len(merged)%2 == 1 {
//		return float64(merged[mid])
//	} else {
//		return (float64(merged[mid+1]) + float64(merged[mid])) / 2.0
//	}
//
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		partitionA := (left + right) / 2
		partitionB := (m+n+1)/2 - partitionA

		maxLeftA := math.MinInt64
		if partitionA != 0 {
			maxLeftA = nums1[partitionA-1]
		}
		minRightA := math.MaxInt64
		if partitionA != m {
			minRightA = nums1[partitionA]
		}

		maxLeftB := math.MinInt64
		if partitionB != 0 {
			maxLeftB = nums2[partitionB-1]
		}
		minRightB := math.MaxInt64
		if partitionB != n {
			minRightB = nums2[partitionB]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			if (m+n)%2 == 0 {
				maxLeft := math.Max(float64(maxLeftA), float64(maxLeftB))
				minRight := math.Min(float64(minRightA), float64(minRightB))
				return (maxLeft + minRight) / 2.0
			} else {
				return math.Max(float64(maxLeftA), float64(maxLeftB))
			}
		} else if maxLeftA > minRightB {
			right = partitionA - 1
		} else {
			left = partitionA + 1
		}
	}
	return 0.0
}
