type MyQueue struct {
    s1, s2 []int
}


func Constructor() MyQueue {
    return MyQueue{
        s1: []int{},
        s2: []int{},
    }
}


func (this *MyQueue) Push(x int)  {
    this.s1 = append(this.s1, x)
}


func(this *MyQueue) flipOneIntoTwo() {
    if len(this.s2) > 0 {return}
    for len(this.s1) != 0 {
        top := this.s1[len(this.s1)-1]
        this.s1 = this.s1[:len(this.s1)-1]
        this.s2 = append(this.s2, top)
    }
}

func (this *MyQueue) Pop() int {
    this.flipOneIntoTwo()    
    top := this.s2[len(this.s2)-1]
    this.s2 = this.s2[:len(this.s2)-1]
    return top
}


func (this *MyQueue) Peek() int {
    this.flipOneIntoTwo()
    return this.s2[len(this.s2)-1]
}


func (this *MyQueue) Empty() bool {
    return len(this.s1) == 0 && len(this.s2) == 0    
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */