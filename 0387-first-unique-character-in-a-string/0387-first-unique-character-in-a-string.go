func firstUniqChar(s string) int {
	hashMap := [26]int{}
	for _, val := range []byte(s) {
		hashMap[val-'a']++
	}
	for i, val := range []byte(s) {
		if hashMap[val-'a'] == 1 {
			return i
		}
	}
	return -1
}