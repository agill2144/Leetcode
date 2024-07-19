type MyCircularQueue struct {
    reader int
    writer int
    data []int
    k int
    count int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        reader: 0,
        writer: 0,
        data: make([]int, k),
        k: k,
        count: 0,
    }
}



func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {return false}
    this.data[this.writer % this.k] = value
    this.count++
    this.writer++  
    return true  
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {return false}
    this.reader++
    this.count--
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {return -1}
    return this.data[this.reader % this.k]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {return -1}
    return this.data[(this.writer-1) % this.k]
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