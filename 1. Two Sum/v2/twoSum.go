package v2

func twoSum(nums []int, target int) []int {
	prevMap := map[int]int{}

	for i, v := range nums {
		require := target - v
		if index, ok := prevMap[require]; ok {

			return []int{i, index}
		}
		prevMap[v] = i
	}

	return nil
}
