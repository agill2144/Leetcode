type MinStack struct {
    st1 []int
    st2 []int // mins
}


func Constructor() MinStack {
    return MinStack{
        st1: []int{},
        st2: []int{},
    }
}


func (this *MinStack) Push(val int)  {
    this.st1 = append(this.st1, val)
    if len(this.st2) == 0 || val <= this.st2[len(this.st2)-1] {
        this.st2 = append(this.st2, val)
    }
}


func (this *MinStack) Pop()  {
    top := this.st1[len(this.st1)-1]
    this.st1 = this.st1[:len(this.st1)-1]
    if len(this.st2) > 0 && top == this.st2[len(this.st2)-1] {
        this.st2 = this.st2[:len(this.st2)-1]
    }
}


func (this *MinStack) Top() int {
    return this.st1[len(this.st1)-1]
}


func (this *MinStack) GetMin() int {
    return this.st2[len(this.st2)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */