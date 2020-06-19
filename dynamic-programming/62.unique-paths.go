package dynamic_programming

func uniquePaths(m int, n int) int {
	res := make([][]int, m, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				res[i][j] = 1
				continue
			}
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}

func uniquePathsSolution2(m int, n int) int {
	res := make([]int, n, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				res[j] = 1
				continue
			}
			res[j] += res[j-1]
		}
	}
	return res[n-1]
}
