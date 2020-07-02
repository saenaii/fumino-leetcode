package dynamic_programming

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var a, b, first, second int
	for i := 0; i < len(nums)-1; i++ {
		first = max(first, max(b, a+nums[i]))
		a, b = b, first
	}

	a, b = 0, 0
	for i := 1; i < len(nums); i++ {
		second = max(second, max(b, a+nums[i]))
		a, b = b, second
	}
	return max(first, second)
}
