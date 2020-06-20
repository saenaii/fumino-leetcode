package dynamic_programming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	res := make([]int, len(obstacleGrid[0]))
	res[0] = 1
	if obstacleGrid[0][0] == 1 {
		res[0] = 0
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				res[j] = 0
				continue
			}
			if j != 0 {
				res[j] += res[j-1]
			}
		}
	}
	return res[len(res)-1]
}
