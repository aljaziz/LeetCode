func intersect(nums1 []int, nums2 []int) []int {
	hashMap := map[int]int{}

	for _, num := range nums1 {
		hashMap[num]++
	}

	ans := []int{}

	for _, num := range nums2 {
		if hashMap[num] > 0 {
			ans = append(ans, num)
			hashMap[num]--
		}
	}
	return ans
}