type MovingAverage struct {
    nums []int
    left int
    sum int
    k int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{nums: []int{}, k: size}    
}


func (this *MovingAverage) Next(val int) float64 {
    this.nums = append(this.nums, val)
    this.sum += val
    right := len(this.nums)-1
    size := right-this.left+1
    if size > this.k {
        this.sum -= this.nums[this.left]
        this.left++
        size = right-this.left+1
    }
    return float64(this.sum)/float64(size)

}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */