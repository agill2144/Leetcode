/*
    push(1)
    push(2)
    push(3)
    
    in a queue = [1,2,3]
    
    in a stack = [1,2,3]
    
    so pushes are identical
    
    pop() - in stack removes top, in queue removes first
    which means inorder to return 1 , we need to flip the stack into a new stack
    therefore we need 2 stacks
    
*/
type MyQueue struct {
    st1 []int
    st2 []int
}


func Constructor() MyQueue {
    return MyQueue{
        st1: []int{},
        st2: []int{},
    }
}


// time ; o(1)
// space ; o(1)
func (this *MyQueue) Push(x int)  {
    this.st1 = append(this.st1, x)
}

// time ; o(n) ; worst case
// space ; o(1)
func(this *MyQueue) flipSt1IntoSt2() {
    if len(this.st2) > 0 {return}
    for len(this.st1) != 0 {
        top := this.st1[len(this.st1)-1]
        this.st1 = this.st1[:len(this.st1)-1]
        this.st2 = append(this.st2, top)
    }
}


// time ; o(n) ; worst case if we need to flip
// space ; o(1)
func (this *MyQueue) Pop() int {
    this.flipSt1IntoSt2()
    top := this.st2[len(this.st2)-1]
    this.st2 = this.st2[:len(this.st2)-1]
    return top
}

// time ; o(n) ; worst case if we need to flip
// space ; o(1)
func (this *MyQueue) Peek() int {
    this.flipSt1IntoSt2()
    return this.st2[len(this.st2)-1]
}


func (this *MyQueue) Empty() bool {
    return len(this.st1) == 0 && len(this.st2) == 0 
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */