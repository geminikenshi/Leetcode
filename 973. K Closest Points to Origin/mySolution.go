package _73__K_Closest_Points_to_Origin

func kClosest(points [][]int, k int) [][]int {
	ans := [][]int{}
	min := 0
	dis := 0
	index := 0
	for (k>0){
		k--
		for i, v := range points{
			if i ==0{
				min = v[0]*v[0] + v[1]*v[1]
			}
			dis = v[0]*v[0] + v[1]*v[1]
			if dis <= min {
				min = dis
				index = i
			}
		}
		ans = append(ans, points[index])
		points = append(points[:index],points[index+1:]...)
	}

	return ans
}