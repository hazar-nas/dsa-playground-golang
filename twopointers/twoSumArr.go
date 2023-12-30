package twopointers

/*
using left and right pointer, trying to find  two elements
yielding target value when adding them up.

input= [1,4,5,7], target 12 should return [3,4]
*/

func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	tempValue := 0

	for {
		tempValue = numbers[left] + numbers[right]

		if left >= right {
			break
		}

		if tempValue > target {
			right--
		} else if tempValue < target {
			left++
		} else if tempValue == target {
			return []int{left + 1, right + 1}
		}

	}

	return []int{0, 0}

}
