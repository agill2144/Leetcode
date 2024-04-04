type stackNode struct {
    val int
    min int
}

type MinStack struct {
    st []stackNode
}


func Constructor() MinStack {
    return MinStack{st: []stackNode{}}    
}


func (this *MinStack) Push(val int)  {
    stNode := stackNode{val: val, min: val}
    if len(this.st) == 0 {
        this.st = append(this.st, stNode)
        return
    }
    peek := this.st[len(this.st)-1]
    if peek.min < stNode.min {
        stNode.min = peek.min
    }
    this.st = append(this.st, stNode)
}


func (this *MinStack) Pop()  {
    this.st = this.st[:len(this.st)-1]    
}


func (this *MinStack) Top() int {
    return this.st[len(this.st)-1].val
}


func (this *MinStack) GetMin() int {
    return this.st[len(this.st)-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */