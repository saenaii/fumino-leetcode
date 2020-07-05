package dynamic_programming

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	res := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i][j-1], min(dp[i-1][j-1], dp[i-1][j])) + 1
			}
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

func maximalSquareSolution2(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}
	dp := make([]int, len(matrix[0]))
	res, previous := 0, 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			temp := dp[j]
			if j == 0 {
				dp[j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				dp[j] = min(previous, min(dp[j-1], dp[j])) + 1
			} else {
				dp[j] = 0
			}
			previous = temp
			res = max(res, dp[j])
		}
	}
	return res * res
}
