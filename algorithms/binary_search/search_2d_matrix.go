package binary_search

/*
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.



Example 1:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
Example 2:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, (len(matrix))*(len(matrix[0]))-1
	rows := len(matrix[0]) // 4
	for l <= r {           //
		mid := (l + r) / 2      // 6 // 2
		col := mid % rows       // 2 //
		row := mid / rows       // 2
		num := matrix[row][col] // 11

		if num == target {
			return true
		}

		if num < target {
			l = mid + 1
		} else {
			r = mid - 1 // r = 5
		}
	}

	return false
}
