package twopointers

import "strings"

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)

	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true

}

func removeNonAlphanumeric(s string) string {
	var result strings.Builder

	for _, char := range s {
		if isAlphanumeric(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}
