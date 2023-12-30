package main

import (
	"dsa-with-go/twopointers"
	"fmt"
)

func main() {
	// var result = arraysandhashing.IsAnagram("silent", "listen")
	// var result = arraysandhashing.TwoSum({}int{2, 7, 11, 15}, 17)
	// var result = arraysandhashing.IsValidSudoku([][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	// })

	// var result = twopointers.IsPalindrome("mIra aRim")
	var result = twopointers.TwoSum([]int{1, 2, 4, 8, 12}, 20)

	fmt.Println(result)
}
