type MovingAverage struct {
    items []int
    sum int
    left int
    right int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        items: make([]int, size),
        sum: 0,
        left: 0,
        right: 0,
    }
}


func (this *MovingAverage) Next(val int) float64 {
    n := len(this.items)
    if this.right-this.left == n {
        this.sum -= this.items[this.left%n]
        this.left++
    }
    this.items[this.right%n] = val
    this.sum += val
    this.right++
    return float64(this.sum) / float64(this.right-this.left)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */