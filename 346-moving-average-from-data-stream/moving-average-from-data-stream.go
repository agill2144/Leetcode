type MovingAverage struct {
    k int
    sum int
    left int
    right int
    nums []int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        sum:0,
        k:size,
        left:0,
        right:0,
        nums: make([]int, size),
    }    
}


func (this *MovingAverage) Next(val int) float64 {
    // right ptr is actually sitting at next cell
    // its not at the cell that it had last written data to
    // therefore we do not have the +1 in this window calc
    if this.right - this.left == this.k {
        this.sum -= this.nums[this.left%this.k]
        this.left++
    }
    this.nums[this.right%this.k] = val
    this.sum += val
    this.right++
    currSize := this.right - this.left
    return float64(this.sum) / float64(currSize)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */