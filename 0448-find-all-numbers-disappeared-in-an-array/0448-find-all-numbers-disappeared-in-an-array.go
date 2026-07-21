func findDisappearedNumbers(nums []int) []int {
	for _, val := range nums {
		val = int(math.Abs(float64(val)))
		nums[val-1] = -int(math.Abs(float64(nums[val-1])))
	}
	ans := []int{}
	for i, val := range nums {
		if val > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}