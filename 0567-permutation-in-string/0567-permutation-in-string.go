import "maps"

func match(s1, s2 map[byte]int) bool {
	return maps.Equal(s1, s2)
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	hashMap1 := map[byte]int{}
	hashMap2 := map[byte]int{}

	for i := 0; i < len(s1); i++ {
		hashMap1[s1[i]]++
		hashMap2[s2[i]]++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if match(hashMap1, hashMap2) {
			return true
		}
		if hashMap2[s2[i]] > 1 {
			hashMap2[s2[i]]--
		} else {
			delete(hashMap2, s2[i])
		}
		hashMap2[s2[i+len(s1)]]++
	}
	return match(hashMap1, hashMap2)
}