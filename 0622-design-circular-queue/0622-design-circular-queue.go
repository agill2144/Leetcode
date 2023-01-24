type MyCircularQueue struct {
    rPtr int
    wPtr int
    count int
    k int
    items []int
}


func Constructor(k int) MyCircularQueue {
    items := make([]int, k)
    return MyCircularQueue{
        k:k,
        items: items,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {return false}
    this.items[this.wPtr % this.k] = value
    this.wPtr++
    this.count++
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {return false}
    _ = this.items[this.rPtr % this.k]
    this.rPtr++
    this.count--
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {return -1}
    return this.items[this.rPtr % this.k]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {return -1}
    return this.items[(this.wPtr-1) % this.k]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.count == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.count == this.k
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