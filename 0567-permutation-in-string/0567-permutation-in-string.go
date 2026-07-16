func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	hashMap1 := [26]int{0}
	hashMap2 := [26]int{0}

	for i := 0; i < len(s1); i++ {
		hashMap1[s1[i]-'a']++
		hashMap2[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if hashMap1 == hashMap2 {
			return true
		}
		hashMap2[s2[i]-'a']--
		hashMap2[s2[i+len(s1)]-'a']++
	}
	return hashMap1 == hashMap2
}