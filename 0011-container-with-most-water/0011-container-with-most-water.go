func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l <= r {
		area := (r - l) * min(height[l], height[r])
		if ans < area {
			ans = area
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}