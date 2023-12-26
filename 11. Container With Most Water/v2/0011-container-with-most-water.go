func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		maxArea = max(maxArea, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
