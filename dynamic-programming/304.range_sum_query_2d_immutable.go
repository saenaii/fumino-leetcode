package dynamic_programming

type NumMatrix struct {
	sum [][]int
}


func ConstructorNumMatrix(matrix [][]int) NumMatrix {
	sum := matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				sum[i][j] = sum[i][j-1] + matrix[i][j]
				continue
			}
			if j == 0 {
				sum[i][j] = sum[i-1][j] + matrix[i][j]
				continue
			}
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{sum: sum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 {
		return this.sum[row2][col2]
	}
	if row1 == 0 {
		return this.sum[row2][col2] - this.sum[row2][col1-1]
	}
	if col1 == 0 {
		return this.sum[row2][col2] - this.sum[row1-1][col2]
	}
	return this.sum[row2][col2] - this.sum[row2][col1-1] - this.sum[row1-1][col2] + this.sum[row1-1][col1-1]
}

// Solution2
func ConstructorNumMatrix2(matrix [][]int) NumMatrix {
	return NumMatrix{sum: matrix}
}

func (this *NumMatrix) SumRegion2(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for i:=row1; i <= row2; i++ {
		for j:=col1; j <= col2; j++ {
			result += this.sum[i][j]
		}
	}
	return result
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
