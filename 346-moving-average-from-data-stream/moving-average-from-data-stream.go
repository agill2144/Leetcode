type MovingAverage struct {
    nums []int
    left int
    right int
    sum int    
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        sum: 0,
        left: 0,
        right: 0,
        nums: make([]int, size),
    }
}


func (this *MovingAverage) Next(val int) float64 {
    n := len(this.nums)
    if this.right - this.left == n {
        this.sum -= this.nums[this.left%n]
        this.left++
    }
    this.sum += val
    this.nums[this.right%n] = val
    this.right++
    return float64(this.sum) / float64(this.right-this.left)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */