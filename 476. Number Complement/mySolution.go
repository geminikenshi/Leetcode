package _76__Number_Complement

import "math"

func findComplement(num int) int {
	var  ans = 0
	var m = 0.0
	if( num > 1){
		if (num == 2){
			ans = num ^ 3
			return ans
		}
		f := float64(num)

		n :=  math.Ceil(math.Log2(f))
		if(n != math.Log2(f)){
			m =  math.Pow(2,n)

		}else{
			m =  math.Pow(2,n+1)

		}

		ans = (int(m)-1) ^ num

	}else{
		ans = num ^ 1
	}

	return ans
}