package v1

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
