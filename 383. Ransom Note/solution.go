package _83__Ransom_Note

func canConstruct(ransomNote string, magazine string) bool {
	rmap := make(map[rune]int)
	mmap := make(map[rune]int)
	for _, v := range magazine{
		if _,ok := mmap[v] ;ok{
			mmap[v] ++
		}else{
			mmap[v] = 1
		}
	}
	for _, v := range ransomNote{
		if _,ok := rmap[v] ;ok{
			rmap[v]++
			if rmap[v] > mmap[v]{
				return false
			}
		}else{
			if mmap[v] == 0{
				return false
			}
			rmap[v] = 1
		}
	}
	// for key,_ := range rmap{             // first time code
	//     if rmap[key] > mmap[key]{
	//         return false
	//     }
	// }

	return true
}