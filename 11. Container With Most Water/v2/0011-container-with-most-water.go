func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		res = max(res, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}