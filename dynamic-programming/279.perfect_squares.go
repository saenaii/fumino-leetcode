package dynamic_programming

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		tmp := n
		for j := 1; j*j <= i; j++ {
			tmp = min(tmp, dp[i-j*j]+1)
		}
		dp[i] = tmp
	}
	return dp[n]
}

func numSquaresSolution2(n int) int {
	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}
	for a := 0; a*a <= n; a++ {
		b := int(math.Sqrt(float64(n - a*a)))
		if a*a+b*b == n {
			if a > 0 && b > 0 {
				return 2
			}
			if a > 0 || b > 0 {
				return 1
			}
		}
	}
	return 3
}
