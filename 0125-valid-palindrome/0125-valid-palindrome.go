func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlphaNumeric(s[l]) {
			l++
		}
		for l < r && !isAlphaNumeric(s[r]) {
			r--
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNumeric(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}