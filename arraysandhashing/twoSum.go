package arraysandhashing

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

//	{

//		2: 0,
//		7: 1,
//		11: 2,
//		15: 3

// }

func TwoSum(nums []int, target int) []int {

	complementMap := make(map[int]int)

	for i, num := range nums {

		complementVal := target - num

		if index, found := complementMap[complementVal]; found {
			return []int{index, i}
		}

		complementMap[num] = i

	}

	return []int{}

}
