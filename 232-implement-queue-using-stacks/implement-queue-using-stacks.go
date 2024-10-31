type MyQueue struct {
    st []int
    st2 []int
}


func Constructor() MyQueue {
    return MyQueue{st: []int{}, st2: []int{}}
}


func (this *MyQueue) Push(x int)  {
    this.st = append(this.st, x)
}

func (this *MyQueue) flip() {
    if len(this.st2) > 0 {return}
    for len(this.st) != 0 {
        top := this.st[len(this.st)-1]
        this.st = this.st[:len(this.st)-1]
        this.st2 = append(this.st2, top)
    }
}


func (this *MyQueue) Pop() int {
    this.flip()
    top := this.st2[len(this.st2)-1]
    this.st2 = this.st2[:len(this.st2)-1]
    return top
}


func (this *MyQueue) Peek() int {
    this.flip()
    return this.st2[len(this.st2)-1]        
}


func (this *MyQueue) Empty() bool {
    return len(this.st) == 0 && len(this.st2) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */