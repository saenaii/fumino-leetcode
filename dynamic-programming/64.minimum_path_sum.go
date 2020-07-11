package dynamic_programming

func minPathSum(grid [][]int) int {
	res := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 {
				res[j] += grid[i][j]
				continue
			}
			if i == 0 {
				res[j] = res[j-1] + grid[i][j]
				continue
			}
			res[j] = grid[i][j] + min(res[j-1], res[j])
		}
	}
	return res[len(res)-1]
}
