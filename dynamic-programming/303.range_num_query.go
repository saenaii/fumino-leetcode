package dynamic_programming

type NumArray struct {
	sum []int
}


func Constructor(nums []int) NumArray {
	obj := NumArray{sum: nums}
	for i := 1; i < len(nums); i++ {
		obj.sum[i] += nums[i-1]
	}
	return obj
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sum[j]
	}
	return this.sum[j] - this.sum[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */