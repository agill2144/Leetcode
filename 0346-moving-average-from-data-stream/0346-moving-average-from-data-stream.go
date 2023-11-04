type MovingAverage struct {
    win []int
    maxSize int
    sum int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{
        win: []int{},
        maxSize: size,
        sum: 0,
    }
}


func (this *MovingAverage) Next(val int) float64 {
    if len(this.win) >= this.maxSize {
        top := this.win[0]
        this.win = this.win[1:]
        this.sum -= top
    }
    this.win = append(this.win, val)
    this.sum += val
    return float64(this.sum) / float64(len(this.win))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */