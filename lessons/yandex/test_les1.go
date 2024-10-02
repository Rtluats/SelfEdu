package main

import "fmt"

// input : []int{1,2,3,4,6,8,9,11,12,13,14}
// output: []string{"1-4","6","8-9","11-14"}

func f(input []int) []string {
	var ans []string
	l, r := 0, 0

	for r < len(input) {

		for r < len(input)-1 && input[r]+1 == input[r+1] {
			r++
		}

		if l == r {
			ans = append(ans, fmt.Sprintf("%v", input[l]))
		} else {
			ans = append(ans, fmt.Sprintf("%v-%v", input[l], input[r]))
		}

		r++
		l = r

	}

	return ans
}

func main() {
	input := []int{1, 2, 3, 4, 6, 8, 9, 11, 12, 13, 14, 17}
	fmt.Println(f(input))
}
