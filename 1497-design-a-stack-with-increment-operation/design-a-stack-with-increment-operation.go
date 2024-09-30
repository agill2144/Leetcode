type CustomStack struct {
    st []int
    maxSize int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{st: []int{}, maxSize: maxSize}    
}

func (this *CustomStack) isFull() bool {
    return len(this.st) == this.maxSize
}

func (this *CustomStack) isEmpty() bool {
    return len(this.st) == 0
}


func (this *CustomStack) Push(x int)  {
    if this.isFull() {return}
    this.st = append(this.st, x) 
}


func (this *CustomStack) Pop() int {
    if this.isEmpty() {return -1}
    out := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    return out
}


func (this *CustomStack) Increment(k int, val int)  {
    for i := 0; i < k && i < len(this.st); i++ {
        this.st[i] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */