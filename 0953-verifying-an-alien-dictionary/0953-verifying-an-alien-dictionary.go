func isAlienSorted(words []string, order string) bool {
	hashMap := map[byte]int{}

	for i := 0; i < len(order); i++ {
		hashMap[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}
			if words[i][j] != words[i+1][j] {
				currLetter := hashMap[words[i][j]]
				nextLetter := hashMap[words[i+1][j]]
				if currLetter > nextLetter {
					return false
				} else {
					break
				}
			}
		}
	}
	return true
}

