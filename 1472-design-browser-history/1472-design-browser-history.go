type node struct {
    val string
    next *node
    prev *node
}
type BrowserHistory struct {
    head *node
    tail *node
    curr *node
}


func Constructor(homepage string) BrowserHistory {
    newNode := &node{val: homepage}
    return BrowserHistory{
        head: newNode,
        tail: newNode,
        curr: newNode,
    }
}


func (this *BrowserHistory) Visit(url string)  {
    newNode := &node{val: url}
    // 1-2 3-4
    //   c   t
    tmp := this.curr.next
    if tmp != nil && tmp.prev != nil {
        tmp.prev = nil
    }
    this.curr.next = nil
    this.tail = this.curr
    
    newNode.prev = this.tail
    this.tail.next = newNode
    this.tail = newNode
    this.curr = newNode
}


func (this *BrowserHistory) Back(steps int) string {
    currPtr := this.curr
    for currPtr.prev != nil && steps != 0 {
        currPtr = currPtr.prev
        steps--
    }
    this.curr = currPtr
    return currPtr.val
}


func (this *BrowserHistory) Forward(steps int) string {
    currPtr := this.curr
    for currPtr.next != nil && steps != 0 {
        currPtr = currPtr.next
        steps--
    }
    this.curr = currPtr
    return currPtr.val
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */