package arraysandhashing

// {
// 	h:1,
// 	e:2,
// 	l:2,
// 	o:1
// }

func IsAnagram(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	charCount := make(map[rune]int)

	for _, char := range str1 {
		charCount[char]++
	}

	for _, char := range str2 {

		if charCount[char] == 0 {
			return false
		}

		charCount[char]--
	}

	return true
}
