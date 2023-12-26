func threeSum(nums []int) [][]int {
	// sort the input array
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	res := [][]int{}

	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if n+nums[l]+nums[r] < 0 {
				l++
			} else if n+nums[l]+nums[r] > 0 {
				r--
			} else {
				res = append(res, []int{n, nums[l], nums[r]})
				// update the pointer
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return res
}