type listNode struct {
    val int
    next *listNode
    prev *listNode
}
type MyCircularDeque struct {
    head *listNode
    tail *listNode
    maxSize int
    size int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        maxSize: k,
    }
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.size == this.maxSize {
        return false
    }
    newNode := &listNode{val: value}
    if this.head == nil {
        this.head = newNode
        this.tail = newNode
        this.size = 1
        return true
    }
    
    newNode.next = this.head
    this.head.prev = newNode
    this.head = newNode
    this.size++
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.size == this.maxSize {
        return false
    }
    newNode := &listNode{val: value}
    if this.head == nil {
        this.head = newNode
        this.tail = newNode
        this.size = 1
        return true
    }
    newNode.prev = this.tail
    this.tail.next = newNode
    this.tail = newNode
    this.size++
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.head == nil {
        return false
    }
    if this.head.next == nil {
        this.head = nil
        this.tail = nil
        this.size = 0
        return true
    }
    newHead := this.head.next
    newHead.prev = nil
    this.head.next = nil
    this.head = newHead
    this.size--
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.head == nil {
        return false
    }
    if this.head.next == nil {
        this.head = nil
        this.tail = nil
        this.size = 0
        return true
    }
    newTail := this.tail.prev
    newTail.next = nil
    this.tail.prev = nil
    this.tail = newTail
    this.size--
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.head == nil {return -1}
    return this.head.val
}


func (this *MyCircularDeque) GetRear() int {
    if this.head == nil {return -1}
    return this.tail.val    
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.size == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return this.size == this.maxSize
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */