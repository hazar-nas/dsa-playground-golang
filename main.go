package main

import (
	"dsa-with-go/arraysandhashing"
	"fmt"
)

func main() {
	// var result = arraysandhashing.IsAnagram("silent", "listen")
	// var result = arraysandhashing.TwoSum({}int{2, 7, 11, 15}, 17)
	var result = arraysandhashing.IsValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})

	// s1 := make([]map[byte]bool, 9)

	// for i := 0; i < 9; i++ {
	// 	s1[i] = make(map[byte]bool)
	// }

	// s1[0][0] = true
	// s1[1][8] = true
	// s1[1][7] = true
	// s1[8][8] = true

	// fmt.Println(s1)

	fmt.Println(result)
}
