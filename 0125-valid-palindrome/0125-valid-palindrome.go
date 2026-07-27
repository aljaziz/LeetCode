import "unicode"

func isPalindrome(s string) bool {
	var newString []rune

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			newString = append(newString, unicode.ToLower(char))
		}
	}
	l, r := 0, len(newString)-1
	for l < r {
		if newString[l] != newString[r] {
			return false
		}
		l++
		r--
	}
	return true
}