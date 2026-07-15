func firstMissingPositive(nums []int) int {
	answerWindow := len(nums)

	foundOne := false
	for i := 0; i < answerWindow; i++ {
		if nums[i] == 1 {
			foundOne = true
		}
		if nums[i] <= 0 || nums[i] > answerWindow {
			nums[i] = 1
		}
	}
	if !foundOne {
		return 1
	}

	for i := 0; i < answerWindow; i++ {
		val := int(math.Abs(float64(nums[i])))
		if val == answerWindow {
			nums[0] = -int(math.Abs(float64(nums[0])))
		} else {
			nums[val] = -int(math.Abs(float64(nums[val])))
		}
	}

	for i := 1; i < answerWindow; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return answerWindow
	}
	return answerWindow + 1
}