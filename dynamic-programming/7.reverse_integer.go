package dynamic_programming

import "math"

func reverse(x int) int {
	if x > 0 {
		return reverseNumber(x)
	}
	return -reverseNumber(-x)
}

func reverseNumber(x int) int {
	res := 0
	a, b := x/10, x%10
	for a > 0 {
		res = (res + b) * 10
		a, b = a/10, a%10
	}
	res += b

	if res > math.MaxInt32 {
		return 0
	}
	return res
}
