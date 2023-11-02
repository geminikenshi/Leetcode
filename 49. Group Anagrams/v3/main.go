package v3

import "sort"

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	anagrams := map[string][]string{}
	ans := [][]string{}
	for _, str := range strs {
		sorted := sortString(str)

		if xs, ok := anagrams[sorted]; !ok {
			anagrams[sorted] = []string{str}
		} else {
			anagrams[sorted] = append(xs, str)
		}
	}

	// fmt.Println(anagrams)

	for _, xs := range anagrams {
		ans = append(ans, xs)
	}

	return ans
}

func sortString(str string) string {

	xrune := []rune(str)
	sort.Slice(xrune, func(i, j int) bool { return xrune[i] < xrune[j] })

	// fmt.Println(sorted)
	return string(xrune)
}
