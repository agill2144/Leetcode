type MovingAverage struct {
    k int
    sum int
    left int
    right int
    nums []int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        k: size,
        nums: make([]int, size),
    }
}


func (this *MovingAverage) Next(val int) float64 {
    if this.right-this.left == this.k {
        this.sum -= this.nums[this.left % this.k]
        this.left++
    }
    this.sum += val
    this.nums[this.right%this.k] = val
    this.right++
    currSize := this.right - this.left
    return float64(this.sum) / float64(currSize)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */