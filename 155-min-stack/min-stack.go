type node struct {
    val int
    min int
}
type MinStack struct {
    st []*node
}


func Constructor() MinStack {
    return MinStack{
        st: []*node{},
    }    
}


func (this *MinStack) Push(val int)  {
    newNode := &node{val:val,min:val}
    if len(this.st) != 0 {
        newNode.min = min(newNode.min, this.st[len(this.st)-1].min)
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