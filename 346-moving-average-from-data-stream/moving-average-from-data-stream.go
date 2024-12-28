type MovingAverage struct {
    nums []int
    left int
    right int
    sum int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{nums: make([]int, size)}    
}


func (this *MovingAverage) Next(val int) float64 {
    k := len(this.nums)
    if this.right-this.left == k {
        this.sum -= this.nums[this.left%k]
        this.left++
    }
    this.sum += val
    this.nums[this.right%k] = val
    this.right++
    return float64(this.sum) / float64(this.right-this.left)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */