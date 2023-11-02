package v2

import "sort"

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	sorted := sortStrings(strs)
	anagrams := map[string][]string{}
	ans := [][]string{}
	for i, ss := range sorted {
		if xs, ok := anagrams[ss]; !ok {
			anagrams[ss] = []string{strs[i]}
		} else {
			anagrams[ss] = append(xs, strs[i])
		}

	}

	// fmt.Println(anagrams)

	for _, xs := range anagrams {
		ans = append(ans, xs)
	}

	return ans
}

func sortStrings(strs []string) []string {
	sorted := []string{}
	for _, str := range strs {
		xrune := []rune(str)
		sort.Slice(xrune, func(i, j int) bool { return xrune[i] < xrune[j] })
		sorted = append(sorted, string(xrune))
	}

	// fmt.Println(sorted)
	return sorted
}
