package _87__First_Unique_Character_in_a_String
func firstUniqChar(s string) int {
	if len(s) == 1{
		return 0
	}
	m := make(map[byte]int)
	for i := 0;i<len(s);i++{        //store how many times a character appears into the map
		if _, ok :=m[s[i]]; ok{     //check if the character is in the map
			m[s[i]]++


		}else{                      //if not, add character into the map
			m[s[i]] = 1

		}
	}
	for i := 0;i <len(s);i++{        //iterate through string again and check the map for value
		if m[s[i]] == 1{             //the first one with only once appearance is the answer
			return i
		}
	}

	return -1

}
