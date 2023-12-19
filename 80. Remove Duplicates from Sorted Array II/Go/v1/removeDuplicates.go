package v1

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	unique := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[unique-2] {
			nums[unique] = nums[i]
			unique++
		}
	}
	return unique
}
