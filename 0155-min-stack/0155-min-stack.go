/*
    approach: brute force
    - getMin() linear scan of stack to find the smallest element
    
    time:
    push() : o(1)
    pop() : o(1)
    getMin() : o(n)
    

    Can we maintian a min in a variable?
    - But then we lose previous min if a latest min gets popped
    - so then maintain another stack that keeps track of all mins
    - when a pop is removing the min element, then we will also remove the top of min stack
    - but what about dupes? if an element is being pushed again and its also the top of the min stack
    - then we will push the same dupe again! so that we can keep same number of mins that are also in the regular stack
 ----------------------------------   
    approach: 2 stacks
    - ceate 2 stack ( regular and min stack )
    - push(x) will push to regular first
        - if the min stack is empty, push there and exit
        - if x is smaller than top of min OR if its the same as top of min ( x <= topOfMinSt )
            - then push this to min stack
    - pop() 
        - pop from top of regular stack
        - if the number being popped is the same as top of the min stack, pop from min stack too
    
    - getMin()
        - return the top of min stack
    time:
    push() : o(1)
    pop() : o(1)
    getMin() : o(n)
    
    space: o(2n) since we used 2 stacks

----------------------------------
    
    can we do this using 1 stack ?
    - yes, we can maintain a pair in regular stack
    - [val, minSoFar]
    - push(x)
        - if this is the first element ever, push it and move on
        - otherwise,
        - create a pair to push [x, x]
        - check with top of regular stack
        - if current min ( top[1] ) is still the min, then update the pair to [x, top[1]]
        - finally push the pair to regular stack
    - pop()
        - blindly pop the top
    - getMin()
        - return top[1]
    time:
    push() : o(1)
    pop() : o(1)
    getMin() : o(n)
    
    space: o(n) 

*/
// type MinStack struct {
//     st []int
//     min []int
// }


// func Constructor() MinStack {
//     return MinStack{
//         st: []int{},
//         min : []int{},
//     }
// }


// func (this *MinStack) Push(val int)  {
//     this.st = append(this.st, val)
//     if len(this.min) == 0 {
//         this.min = append(this.min, val)
//         return
//     }
//     topMin := this.min[len(this.min)-1]
//     if val <= topMin {
//         this.min = append(this.min, val)
//     }
// }


// func (this *MinStack) Pop()  {
//     top := this.st[len(this.st)-1]; this.st = this.st[:len(this.st)-1]
//     if top == this.min[len(this.min)-1] {
//         this.min = this.min[:len(this.min)-1]
//     }
// }


// func (this *MinStack) Top() int {
//     return this.st[len(this.st)-1]
// }


// func (this *MinStack) GetMin() int {
//     return this.min[len(this.min)-1]
// }


type MinStack struct {
    st [][]int // [val, minSoFar]
}


func Constructor() MinStack {
    return MinStack{
        st: [][]int{},
    }
}


/* 
time : o(1)
space: o(1)
*/
func (this *MinStack) Push(val int)  {
    item := []int{val, val}
    if len(this.st) == 0 {
        this.st = append(this.st, item)
        return
    }
    topMin := this.st[len(this.st)-1][1]
    if topMin < val {
        item[1] = topMin
    }
    this.st = append(this.st, item)
}


/* 
time : o(1)
space: o(1)
*/
func (this *MinStack) Pop()  {
    this.st = this.st[:len(this.st)-1]
}


/* 
time : o(1)
space: o(1)
*/
func (this *MinStack) Top() int {
    return this.st[len(this.st)-1][0]
}


/* 
time : o(1)
space: o(1)
*/
func (this *MinStack) GetMin() int {
    return this.st[len(this.st)-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */