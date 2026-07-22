func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashMap := map[int]int{}
	stack := []int{}

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			hashMap[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for i, num := range nums1 {
		if val, ok := hashMap[num]; ok {
			nums1[i] = val
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
