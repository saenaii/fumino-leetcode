package dynamic_programming

func maxProduct(nums []int) int {
	large, small := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		temp := large * nums[i]
		large = max(nums[i], max(temp, small*nums[i]))
		small = min(nums[i], min(temp, small*nums[i]))
		res = max(large, res)
	}
	return res
}
