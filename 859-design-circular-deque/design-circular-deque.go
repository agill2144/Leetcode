type MyCircularDeque struct {
    dq []int
    k int
    f int
    r int
    size int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        dq: make([]int, k),
        k: k,
        f: 0,
        r: k-1,
        size: 0,
    }
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {return false}
    this.f = (this.f-1 + this.k) % this.k
    this.dq[this.f] = value
    this.size++
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {return false}
    this.r = (this.r+1) % this.k
    this.dq[this.r] = value
    this.size++
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {return false}
    this.f = (this.f+1) % this.k
    this.size--
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {return false}
    this.r = (this.r-1+this.k) % this.k
    this.size--
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {return -1}    
    return this.dq[this.f]
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {return -1}    
    return this.dq[this.r]    
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.size == 0
}


func (this *MyCircularDeque) IsFull() bool {
    return this.size == this.k    
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