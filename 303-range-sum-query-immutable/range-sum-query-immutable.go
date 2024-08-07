type NumArray struct {
    prefixSum []int
}


func Constructor(nums []int) NumArray {
    sums := []int{}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        sums = append(sums, rSum)
    }
    return NumArray{prefixSum: sums}
}


func (this *NumArray) SumRange(left int, right int) int {
    x := this.prefixSum[right]
    y := 0
    if left-1 >= 0 {y = this.prefixSum[left-1]}
    return x-y    
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */