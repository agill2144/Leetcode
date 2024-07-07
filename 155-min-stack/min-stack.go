type stNode struct {
    val int
    minval int
}
type MinStack struct {
    st []*stNode
}


func Constructor() MinStack {
    return MinStack{
        st: []*stNode{},
    }
}


func (this *MinStack) Push(val int)  {
    newNode := &stNode{val, val}
    if len(this.st) != 0 && this.st[len(this.st)-1].minval < val {
        newNode.minval = this.st[len(this.st)-1].minval
    }
    this.st = append(this.st, newNode)
}


func (this *MinStack) Pop()  {
    this.st = this.st[:len(this.st)-1]
}


func (this *MinStack) Top() int {
    return this.st[len(this.st)-1].val    
}


func (this *MinStack) GetMin() int {
    return this.st[len(this.st)-1].minval    
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */