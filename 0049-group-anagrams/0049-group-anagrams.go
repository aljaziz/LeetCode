type hashKey [26]byte

func getKey(str string) hashKey {
	var key hashKey
	for i := range str {
		key[str[i]-'a']++
	}
	return key
}
func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }
	hashMap := make(map[hashKey][]string)
	for _, s := range strs {
		key := getKey(s)
		hashMap[key] = append(hashMap[key], s)
	}
	result := make([][]string, 0, len(hashMap))
	for _, i := range hashMap {
		result = append(result, i)
	}
	return result
}