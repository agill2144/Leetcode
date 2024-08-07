type NumArray struct {
    nums []int
    prefix []int
}


func Constructor(nums []int) NumArray {
    sums := []int{}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        sums = append(sums, rSum)        
    }    
    return NumArray{nums:nums,prefix:sums}
}


func (this *NumArray) Update(index int, val int)  {
    currVal := this.nums[index]
    adjustBy := val-currVal
    this.nums[index] = val
    for i := index; i < len(this.prefix); i++ {
        this.prefix[i] += adjustBy
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    x := this.prefix[right]
    y := 0
    if left-1 >= 0 { y = this.prefix[left-1]}
    return x-y
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */