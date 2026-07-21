func intersection(nums1 []int, nums2 []int) []int {
	hashMap1 := map[int]bool{}
	hashMap2 := map[int]bool{}

	for _, num := range nums1 {
		hashMap1[num] = true
	}

	for _, num := range nums2 {
		hashMap2[num] = true
	}
	ans := []int{}
	for num, _ := range hashMap1 {
		if hashMap2[num] {
			ans = append(ans, num)
		}
	}
	return ans
}