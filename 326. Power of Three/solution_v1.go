package _26__Power_of_Three

func isPowerOfThree(n int) bool {
	if n> 0{

		for n != 1{
			if n % 3 != 0{
				return false
			}
			n /= 3
		}
		return true
	}
	return false
}
