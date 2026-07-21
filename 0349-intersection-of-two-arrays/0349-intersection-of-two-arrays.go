func intersection(nums1 []int, nums2 []int) []int {
	hashMap := map[int]bool{}
	for _, num := range nums1 {
		hashMap[num] = true
	}
	ans := []int{}
	for _, num := range nums2 {
		if hashMap[num] {
			ans = append(ans, num)
			hashMap[num] = false
		}
	}
	return ans
}