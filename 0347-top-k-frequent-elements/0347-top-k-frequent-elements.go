type customArr struct {
	key   int
	value int
}

func topKFrequent(nums []int, k int) []int {
	var countArr []customArr
	hashMap := make(map[int]int)
	var ans []int
	for _, num := range nums {
		hashMap[num]++
	}

	for k, v := range hashMap {
		countArr = append(countArr, customArr{k, v})
	}

	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i].value > countArr[j].value
	})

	for i := 0; i < k; i++ {
		ans = append(ans, countArr[i].key)
	}
	return ans
}