package dynamic_programming

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		current := 0
		for j := 0; j < i; j++ {
			if next := dp[j] * 2; next > dp[i-1] {
				current = next
				break
			}
		}
		for j := 0; j < i; j++ {
			if next := dp[j] * 3; next > dp[i-1] {
				current = min(current, next)
				break
			}
		}
		for j := 0; j < i; j++ {
			if next := dp[j] * 5; next > dp[i-1] {
				current = min(current, next)
				break
			}
		}
		dp[i] = current
	}
	return dp[n-1]
}

func nthUglyNumberSolution2(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	var i2, i3, i5 int
	for i := 1; i < n; i++ {
		dp[i] = min(dp[i2]*2, min(dp[i3]*3, dp[i5]*5))
		if dp[i] == dp[i2]*2 {
			i2++
		}
		if dp[i] == dp[i3]*3 {
			i3++
		}
		if dp[i] == dp[i5]*5 {
			i5++
		}
	}
	return dp[n-1]
}
