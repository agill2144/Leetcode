type MyStack struct {
    q []int
}


func Constructor() MyStack {
    return MyStack{q: []int{}}    
}


func (this *MyStack) Push(x int)  {
    this.q = append(this.q, x)
    qSize := len(this.q)
    for qSize != 1 {
        dq := this.q[0]
        this.q = this.q[1:]
        this.q = append(this.q, dq)
        qSize--
    }
}

func (this *MyStack) Pop() int {
    dq := this.q[0]
    this.q = this.q[1:]
    return dq
}


func (this *MyStack) Top() int {
    return this.q[0]
}


func (this *MyStack) Empty() bool {
    return len(this.q) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */