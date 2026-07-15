func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := hashMap[num]; ok {
			if i-j <= k {
				return true
			}
		}
		hashMap[num] = i
	}
	return false
}