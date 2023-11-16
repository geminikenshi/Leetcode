package iterative

func fib(n int) int {
	if n <= 1 {
		return n
	}
	// F(n) = F(n1) + F(n2)
	n2 := 0
	n1 := 1
	fibNumber := 1
	for i := 2; i <= n; i++ {
		fibNumber = n1 + n2
		n2, n1 = n1, fibNumber
	}
	return fibNumber
}
