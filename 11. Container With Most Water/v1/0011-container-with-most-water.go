func maxArea(height []int) int {
	l := 0
	max := 0

	for i, h := range height {
		if h >= height[l] {
			l = i
			r := l + 1
			for r < len(height) {
				if min(height[l], height[r])*(r-l) > max {
					max = min(height[l], height[r]) * (r - l)
				}
				r++
			}
		}
	}

	return max
}

