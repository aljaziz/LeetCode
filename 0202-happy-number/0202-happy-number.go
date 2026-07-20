func squaredSum(n int) int {
	sum := 0
	for n > 0 {
		rem := n % 10
		sum += rem * rem
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	hashMap := map[int]bool{}

	for n != 1 {
		sum := squaredSum(n)
		if sum == 1 {
			return true
		}
		if hashMap[sum] {
			return false
		}
		hashMap[sum] = true
		n = sum
	}
	return true
}