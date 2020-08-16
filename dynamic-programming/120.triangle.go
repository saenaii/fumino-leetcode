package dynamic_programming

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	dp := make([]int, length)
	dp[0] = triangle[0][0]
	for i := 1; i < length; i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			switch j {
			case len(triangle[i]) - 1:
				dp[j] = dp[j-1] + triangle[i][j]
			case 0:
				dp[j] += triangle[i][j]
			default:
				dp[j] = triangle[i][j] + min(dp[j-1], dp[j])
			}
		}
	}

	ans := dp[0]
	for _, v := range dp {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func minimumTotalSolution2(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
