package _338__Reduce_Array_Size_to_The_Half

//This is uncomplete code, I was stuck.

import "sort"

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for i := 0;i<len(arr);i++{
		if _, ok := m[arr[i]]; ok{
			arr[i]++
		}else{
			m[arr[i]] = 1
		}
	}
	length := len(m)
	keys := make([]int, 0, length)
	for k, _ := range m{
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int)bool{  //using ananymous func as parameter
		return m[keys[i]] > m[keys[j]]
	})
	n := length
	for index,value := range keys{
		delete(m, index)
		n -= value
	}



	return 0
}