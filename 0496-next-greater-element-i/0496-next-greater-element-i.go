func getGreaterElement(nums []int, val int) int {
	for _, n := range nums {
		if n > val {
			return n
		}
	}
	return -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashMap := map[int]int{}
	for i, val := range nums2 {
		if i+1 < len(nums2) {
			hashMap[val] = getGreaterElement(nums2[i+1:], val)
		} else {
			hashMap[val] = -1
		}
	}
	for i, val := range nums1 {
		nums1[i] = hashMap[val]
	}
	return nums1
}