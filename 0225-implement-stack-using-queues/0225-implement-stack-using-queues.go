type MyStack struct {
    q1 []int
    q2 []int
}


func Constructor() MyStack {
    return MyStack{q1: []int{}, q2: []int{}}
}


func (this *MyStack) Push(x int)  {
    this.q1 = append(this.q1, x)
}

func (this *MyStack) flip() {
    if len(this.q2) > 0 {return}
    for len(this.q1) != 0 {
        last := this.q1[len(this.q1)-1]
        this.q1 = this.q1[:len(this.q1)-1]
        this.q2 = append(this.q2, last)
    }
}


func (this *MyStack) Pop() int {
    this.flip()
    dq := this.q2[0]
    this.q2 = this.q2[1:]
    return dq
}


func (this *MyStack) Top() int {
    this.flip()
    if len(this.q1) > 0 {
        return this.q1[len(this.q1)-1]
    }
    return this.q2[0]
}


func (this *MyStack) Empty() bool {
    return len(this.q1) == 0 && len(this.q2) == 0
}



/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */