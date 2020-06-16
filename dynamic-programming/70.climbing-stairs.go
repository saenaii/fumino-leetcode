package dynamic_programming

func climbStairs(n int) int {
	first, second := 1, 1
	for i := 1; i < n; i++ {
		first, second = second, first+second
	}
	return second
}
