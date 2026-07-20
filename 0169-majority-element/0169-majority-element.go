func majorityElement(nums []int) int {
	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}

	for i, j := range hashMap {
		if j > len(nums)/2 {
			return i
		}
	}
	return nums[0]
}