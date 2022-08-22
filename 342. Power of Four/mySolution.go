package _42__Power_of_Four

import "math"

func isPowerOfFour(n int) bool {
	if n == 0{
		return false
	}
	i := math.Log(float64(n))/math.Log(float64(4)) //Log4(n)
	if i == math.Floor(i){
		return true
	}
	return false
}