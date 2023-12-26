func maxArea(height []int) int {
	l := 0
	max := 0

	for i, h := range height {
		if h >= height[l] {
			l = i
			r := l + 1
			for r < len(height) {
				if height[r] < height[l] {
					if height[r]*(r-l) > max {
						max = height[r] * (r - l)
					}
				} else {
					if height[l]*(r-l) > max {
						max = height[l] * (r - l)
					}
				}
				r++
			}
		}
	}

	return max
}