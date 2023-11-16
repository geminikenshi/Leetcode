package insertion

func sortArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		j := i - 1
		for j >= 0 && nums[j] > nums[j+1] {
			// swap
			nums[j], nums[j+1] = nums[j+1], nums[j]
			j--

		}
	}

	return nums
}
