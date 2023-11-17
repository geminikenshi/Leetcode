package go

func getConcatenation(nums []int) []int {
    return append(nums, nums...)
}