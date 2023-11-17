package v2

func getConcatenation(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums)*2; i++ {
		ans = append(ans, nums[i%len(nums)])
	}

	return ans
}
