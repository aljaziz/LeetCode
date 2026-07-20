func majorityElement(nums []int) int {
	result, maxOccur := 0, 0

	for _, num := range nums {
		if maxOccur == 0 {
			result = num
		}

		if result == num {
			maxOccur++
		} else {
			maxOccur--
		}
	}
	return result
}