package arrays_and_hashing

import "testing"

//[["8","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]

// |"8"|"3"|"."|"."|"7"|"."|"."|"."|"."|
// |"6"|"."|"."|"1"|"9"|"5"|"."|"."|"."]
//,[".","9","8",".",".",".","."|"6"|"."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]

func getColumns(board [][]byte, column int) []byte {
	res := make([]byte, 0, 9)

	for i := 0; i < 9; i++ {
		res = append(res, board[i][column])
	}

	return res
}

var indexes = []int{0, 1, 2}

func getSquare(board [][]byte, index int) []byte {
	var kc, kr int
	values := make([]byte, 0, 9)
	kr = index % 3

	if index > 2 && index <= 5 {
		kc = 1
	} else if index > 5 {
		kc = 2
	}

	for c := 0 + (kc * 3); c < (kc*3)+3; c++ {
		for _, r := range indexes {
			r += 3 * kr
			values = append(values, board[r][c])
		}
	}

	return values
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ { // 3
		if !validSequence(board[i]) {
			return false
		}

		columns := getColumns(board, i)

		if !validSequence(columns) {
			return false
		}

		square := getSquare(board, i)

		if !validSequence(square) {
			return false
		}
	}

	return true
}

func validSequence(board []byte) bool {
	set := make(map[byte]bool)

	for i := 0; i < len(board); i++ {
		if string(board[i]) == "." {
			continue
		}

		if set[board[i]] {
			return false
		}
		set[board[i]] = true
	}

	return true
}

func Test_fun(t *testing.T) {
	tcs := map[string]struct {
		input    [][]byte
		expected bool
	}{
		"1": {
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: true,
		},
	}

	for n, tc := range tcs {
		t.Run(n, func(t *testing.T) {
			r := isValidSudoku(tc.input)
			if r != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, r)
			}
		})
	}
}
