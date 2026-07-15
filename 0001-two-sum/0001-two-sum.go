func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
    for i, num := range nums {
        rem := target - num
        j, ok := hashMap[rem]
        if ok {
           return []int{i, j}
        } else {
            hashMap[num] = i
        }
    }
    return []int{}
}