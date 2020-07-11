package dynamic_programming

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	length := len(s)
	dp := make([]int, length)
	dp[0] = 1
	for i := 1; i < length; i++ {
		if s[i] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
			if i > 1 {
				dp[i] += dp[i-2]
			} else {
				dp[i]++
			}
		}
	}
	return dp[length-1]
}
