package _82__Baseball_Game

import "strconv"

func calPoints(operations []string) int {
	if len(operations) == 0{
		return 0
	}
	stack := []int{}
	n := 0
	for _, v := range operations{
		n = len(stack)
		switch v{
		case "+":
			stack = append(stack, stack[n-1]+stack[n-2])
		case "C":
			stack = stack[:n-1]
		case "D":
			stack = append(stack, stack[n-1]*2)
		default:
			toInt, _ := strconv.Atoi(v)
			stack = append(stack, toInt)
		}
	}
	n = 0
	for _, v := range stack{
		n += v
	}
	return n
}