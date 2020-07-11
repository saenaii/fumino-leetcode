package dynamic_programming

import "math"

func maxSubArray(nums []int) int {
	sum, res := math.MinInt32, math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}
		res = max(sum, res)
	}
	return res
}
