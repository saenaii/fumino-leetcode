package dynamic_programming

func rob(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	if length <= 2 {
		if length == 1 {
			return nums[0]
		}
		return max(nums[0], nums[1])
	}

	dp := make([]int, length)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[length-1]
}

func robSolution2(nums []int) int {
	var a, b, res int
	for _, v := range nums {
		res = max(res, max(a+v, b))
		a, b = b, res
	}
	return res
}
