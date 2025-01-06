type MyCircularQueue struct {
    count int    
    w int
    r int
    nums []int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        nums: make([]int, k),
    }    
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {return false}
    k := len(this.nums)
    this.nums[this.w%k] = value
    this.count++
    this.w++
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {return false}
    k := len(this.nums)
    this.nums[this.r%k] = 0
    this.count--
    this.r++
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {return -1}
    k := len(this.nums)
    return this.nums[this.r%k]        
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {return -1}
    k := len(this.nums)
    w := (this.w%k)-1
    if w == -1 {w = k-1}
    return this.nums[w]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.count == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.count == len(this.nums)
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */