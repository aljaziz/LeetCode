func canConstruct(ransomNote string, magazine string) bool {
	hashMap := [26]int{}

	for i := 0; i < len(magazine); i++ {
		hashMap[magazine[i]-'a']++
	}

	for _, num := range ransomNote {
		if hashMap[num-'a'] == 0 {
			return false
		}
		hashMap[num-'a']--
	}
	return true
}