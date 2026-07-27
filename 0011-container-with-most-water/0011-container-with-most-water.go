func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		if height[l] < height[r] {
			h := height[l]
			area := h * (r - l)
			if area > ans {
				ans = area
			}
			l++
		} else {
			h := height[r]
			area := h * (r - l)
			if area > ans {
				ans = area
			}
			r--
		}

	}
	return ans
}