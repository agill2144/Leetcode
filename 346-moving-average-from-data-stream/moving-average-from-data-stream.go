type MovingAverage struct {
    nums []int
    left int
    size int
    sum int    
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        nums: []int{},
        size: size,
        left: 0,
        sum: 0,
    }
}


func (this *MovingAverage) Next(val int) float64 {
    this.nums = append(this.nums, val)
    this.sum += val
    if len(this.nums) > this.size {
        this.sum -= this.nums[this.left]
        this.left++
    }
    return float64(this.sum) / float64(len(this.nums)-1 - this.left+1)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */