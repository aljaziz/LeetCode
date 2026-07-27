func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := height[0], height[len(height)-1]
	ans := 0

	for l < r {
		if height[l] < height[r] {
			leftMax = max(leftMax, height[l])
			total := leftMax - height[l]
			if total > 0 {
				ans += total
			}
			l++
		} else {
			rightMax = max(rightMax, height[r])
			total := rightMax - height[r]
			if total > 0 {
				ans += total
			}
			r--
		}
	}
	return ans
}