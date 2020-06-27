package dynamic_programming

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && stringInSlice(wordDict, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
