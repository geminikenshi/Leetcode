package mergesort

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

func mergeSort(nums []int, start, end int) []int {
	if end-start <= 0 {
		return nums
	}
	mid := start + end/2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	return nums
}

func merge(nums []int, start, mid, end int) {
	i, j, k := 0, 0, start
	// left := nums[start : mid+1]
	// right := nums[mid+1 : end+1]
	// This will lead to errors, as new slice from slicing a slice reference to the same underlying array, so any changes to the original slice will affect the new slice as well
	// So, making a copy for each subslice is required

	left := make([]int, mid-start+1)
	right := make([]int, end-mid)

	copy(left, nums[start:mid+1])
	copy(right, nums[mid+1:end+1])

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}
}
