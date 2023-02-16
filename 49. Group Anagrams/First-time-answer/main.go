package firsttimeanswer

import "reflect"

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	ans := [][]string{}
	frequency := []map[rune]int{}
	for _, s := range strs {
		rMap := map[rune]int{}
		for _, r := range s {
			if _, ok := rMap[r]; !ok {
				rMap[r] = 1
			} else {
				rMap[r]++
			}
		}
		//   first string
		if len(frequency) == 0 {
			frequency = append(frequency, rMap)
			first := []string{s}
			ans = append(ans, first)
			continue
		}
		// rest
		equal := false
		for i, f := range frequency {
			if reflect.DeepEqual(f, rMap) {
				ans[i] = append(ans[i], s)
				equal = true
			}
		}
		if !equal {
			frequency = append(frequency, rMap)
			newSlice := []string{s}
			ans = append(ans, newSlice)
		}

	}

	return ans
}
