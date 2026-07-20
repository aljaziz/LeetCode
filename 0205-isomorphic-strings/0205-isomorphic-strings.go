func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap1 := make([]int, 128)
	hashMap2 := make([]int, 128)

	for i := 0; i < len(s); i++ {
		if hashMap1[s[i]] == 0 && hashMap2[t[i]] == 0 {
			hashMap1[s[i]] = int(t[i])
			hashMap2[t[i]] = int(s[i])
		} else if hashMap1[s[i]] != int(t[i]) || hashMap2[t[i]] != int(s[i]) {
			return false
		}
	}
	return true
} 