
/*
    approach:
    
    Since we are given 2 stacks,
    we will use one as the primary pushing stack
    and other to pop/peek from 
    
    But how will pop/peek from 2nd stack work?
    Stack2 should have 1st item at the top ( so items must be in a queue format )
    - Then we will do exactly that
    - As soon as, we hit a pop/peek we will flip all of stack1 into stack2 ( ONLY IF stack2 is empty )
    - the flip will be o(n) time worse case
    
*/

type MyQueue struct {
    stack1 []int
    stack2 []int
}


func Constructor() MyQueue {
    return MyQueue{stack1: []int{}, stack2: []int{}}
}


func (this *MyQueue) flipStack1IntoStack2() {
    if len(this.stack2) == 0 {
        for len(this.stack1) != 0 {
            top := this.stack1[len(this.stack1)-1]
            this.stack1 = this.stack1[:len(this.stack1)-1]
            this.stack2 = append(this.stack2, top)
        }
    }
}

func (this *MyQueue) Push(x int)  {
    this.stack1 = append(this.stack1, x)
}


func (this *MyQueue) Pop() int {
    this.flipStack1IntoStack2()
    out := this.stack2[len(this.stack2)-1]
    this.stack2 = this.stack2[:len(this.stack2)-1]
    return out
}


func (this *MyQueue) Peek() int {
    this.flipStack1IntoStack2()
    return this.stack2[len(this.stack2)-1]
}


func (this *MyQueue) Empty() bool {
    return len(this.stack1) + len(this.stack2) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */