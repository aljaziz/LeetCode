func findDisappearedNumbers(nums []int) []int {
	hashMap := map[int]bool{}
	for _, num := range nums {
		hashMap[num] = true
	}
	ans := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := hashMap[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}