type MovingAverage struct {
    items []int
    sum int
    left int
    size int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        items: []int{},
        sum: 0,
        left: 0,
        size: size,
    }
}


func (this *MovingAverage) Next(val int) float64 {
    this.items = append(this.items, val)
    this.sum += val
    right := len(this.items)-1
    currSize := right-this.left+1
    if currSize > this.size {
        this.sum -= this.items[this.left]
        this.left++
        currSize--
    }
    return float64(this.sum) / float64(currSize)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */