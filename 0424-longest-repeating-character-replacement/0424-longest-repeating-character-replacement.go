func characterReplacement(s string, k int) int {
	hashMap := [26]int{}
	left, right := 0, 0
	ans, maxOccur := 0, 0

	for right < len(s) {
		hashMap[s[right]-'A']++
		if hashMap[s[right]-'A'] > maxOccur {
			maxOccur = hashMap[s[right]-'A']
		}
		currLen := (right - left) + 1
		if k >= (currLen - maxOccur) {
			if currLen > ans {
				ans = currLen
			}
		} else {
			hashMap[s[left]-'A']--
			left++
		}
        right++
	}
	return ans
}