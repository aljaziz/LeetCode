func firstMissingPositive(nums []int) int {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	smallPos := 1
	for i := 0; i < len(nums); i++ {
        if _, ok := hashMap[smallPos]; !ok {
            return smallPos
        }
        smallPos++
	}
    return smallPos
}