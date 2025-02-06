type MovingAverage struct {
    k int
    left int
    sum int
    nums []int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{left:0,nums:[]int{}, k:size, sum:0}    
}


func (this *MovingAverage) Next(val int) float64 {
    this.nums = append(this.nums, val)
    this.sum += val
    right := len(this.nums)-1
    if right-this.left+1 > this.k {
        this.sum -= this.nums[this.left]
        this.left++
    }
    currSize := right-this.left+1
    return float64(this.sum) / float64(currSize)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */