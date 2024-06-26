type MovingAverage struct {
    arr []int
    left int
    sum int
    winsize int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        arr: []int{},
        left: 0,
        sum: 0,
        winsize: size,
    }
}


func (this *MovingAverage) Next(val int) float64 {
    this.sum += val
    this.arr = append(this.arr, val)
    if len(this.arr) > this.winsize {
        leftVal := this.arr[this.left]
        this.sum -= leftVal
        this.left++
    }
    winSize := (len(this.arr)-1) - this.left + 1
    return float64(this.sum) / float64(winSize)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */