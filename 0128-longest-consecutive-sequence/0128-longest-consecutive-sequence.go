func longestConsecutive(nums []int) int {
    if len(nums)==0{
        return 0
    }

	hashSet := make(map[int]bool)

	for _, num := range nums {
		hashSet[num] = true
	}
	ans := 1
	for num := range hashSet {
		if !hashSet[num-1] {
			currNum := num
			currLength := 1
			for hashSet[currNum+1] {
				currNum++
				currLength++
			}
			if currLength > ans {
				ans = currLength
			}
		}
	}
	return ans
}