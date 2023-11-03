package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	frequencyMap := map[int]int{}
	frequencySlice := []frequency{}
	for _, num := range nums {
		if _, ok := frequencyMap[num]; !ok {
			frequencyMap[num] = 1
		} else {
			frequencyMap[num]++
		}
	}
	for number, times := range frequencyMap {
		frequencySlice = append(frequencySlice, frequency{number: number, times: times})
	}
	sort.Slice(frequencySlice, func(i, j int) bool { return frequencySlice[i].times > frequencySlice[j].times })
	for i := 0; i < k; i++ {
		ans = append(ans, frequencySlice[i].number)
	}
	return ans
}

type frequency struct {
	number int
	times  int
}
