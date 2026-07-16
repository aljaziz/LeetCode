func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	pal, temp := 0, x
	for temp > 0 {
		rem := temp % 10
		pal = (pal * 10) + rem
		temp /= 10
	}

	return pal == x

}