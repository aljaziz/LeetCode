func lengthOfLongestSubstring(s string) int {
	hashMap := map[byte]bool{}
	ans, left := 0, 0
	for i := 0; i < len(s); i++ {
		for hashMap[s[i]] {
			delete(hashMap, s[left])
			left++
		}
		hashMap[s[i]] = true

		temp := i - left + 1
		if temp > ans {
			ans = temp
		}
	}
	return ans
}