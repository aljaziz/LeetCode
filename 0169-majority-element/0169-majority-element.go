func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}
	ans := 0

	for i, j := range hashMap {
		if j > len(nums)/2 {
			ans = i
		}
	}
	return ans
}