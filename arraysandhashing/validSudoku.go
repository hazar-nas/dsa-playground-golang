package arraysandhashing

import "fmt"

// [
//
//	map[
//
// 7:true
// 2:true
// 8:true
// 9:true
//
//	....
//	]
//	map[
//
// 5:true
// 0:true
// 1:true
// 4:true
//
//	....
//	]
// ]

func IsValidSudoku(board [][]byte) bool {

	var rows = make([]map[byte]bool, 9)
	var columns = make([]map[byte]bool, 9)
	var boxes = make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for k := 0; k < 9; k++ {
		for j := 0; j < 9; j++ {
			boardValue := board[k][j]

			if boardValue != '.' {

				if rows[k][boardValue] {
					fmt.Println("row index:", k, "breaking rule")
					return false
				}
				rows[k][boardValue] = true

				if columns[j][boardValue] {
					fmt.Println("column index:", j, "breaking rule")

					return false
				}
				columns[j][boardValue] = true

				boardIndex := (k/3)*3 + j/3

				if boxes[boardIndex][boardValue] {
					fmt.Println(boxes[boardIndex][boardValue], "breaking rule")
					return false
				}
				boxes[boardIndex][boardValue] = true

			}
		}
	}

	return true
}

/*
func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for k := 0; k < 9; k++ {
			value := board[i][k]

			if value != '.' {

				if rows[i][value] {
					return false
				}
				rows[i][value] = true

				if columns[k][value] {
					return false
				}
				columns[k][value] = true

				boxIndex := (i/3)*3 + k/3

				if boxes[boxIndex][value] {
					return false
				}
				boxes[boxIndex][value] = true
			}

		}

	}

	return true
}
*/
