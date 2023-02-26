type node struct {
    val int
    next *node
    prev *node
}
type MaxStack struct {
    head *node
    maxH *maxHeap
    removedNodes map[*node]struct{}
}


func Constructor() MaxStack {
    mx := &maxHeap{items: []*node{}}
    return MaxStack{maxH: mx, removedNodes: map[*node]struct{}{}}
}


func (this *MaxStack) Push(x int)  {
    newNode := &node{val: x}
    if this.head == nil {
        this.head = newNode
        heap.Push(this.maxH, newNode)
        return
    }
    newNode.next = this.head
    this.head.prev = newNode
    this.head = newNode
    heap.Push(this.maxH, newNode)
    
}

func (this *MaxStack) adjustStWithRemovedNodes() {
    _, topIsRemovedFromHeap := this.removedNodes[this.head]
    for topIsRemovedFromHeap && this.head != nil {
        newHead := this.head.next
        this.head.next = nil
        newHead.prev = nil
        this.head = newHead
        _, topIsRemovedFromHeap = this.removedNodes[this.head]
    }
}

func (this *MaxStack) adjustHeapWithRemovedNodes() {
    _, topIsRemovedFromSt := this.removedNodes[this.maxH.items[0]]
    for topIsRemovedFromSt && this.maxH.Len() != 0 {
        heap.Pop(this.maxH)
        _, topIsRemovedFromSt = this.removedNodes[this.maxH.items[0]]
    }
}


func (this *MaxStack) Pop() int {
    this.adjustStWithRemovedNodes()
    out := this.head
    if this.head.next == nil {
        this.head = nil
    } else {
        newHead := this.head.next
        newHead.prev = nil
        this.head.next = nil
        this.head = newHead
    }
    this.removedNodes[out] = struct{}{}
    return out.val
}


func (this *MaxStack) Top() int {
    this.adjustStWithRemovedNodes()
    return this.head.val
}


func (this *MaxStack) PeekMax() int {
    this.adjustHeapWithRemovedNodes()
    return this.maxH.items[0].val
}


func (this *MaxStack) PopMax() int {
    this.adjustHeapWithRemovedNodes()
    popped := heap.Pop(this.maxH).(*node)
    this.removedNodes[popped] = struct{}{}
    return popped.val
}


type maxHeap struct {
    items []*node
}
func (mx *maxHeap) Push(x interface{}){
    mx.items = append(mx.items, x.(*node))
}
func (mx *maxHeap) Pop() interface{} {
    out := mx.items[len(mx.items)-1]
    mx.items = mx.items[:len(mx.items)-1]
    return out
}
func (mx *maxHeap) Less(i, j int) bool{
    return mx.items[i].val >= mx.items[j].val
}
func(mx *maxHeap) Swap(i, j int) {
    mx.items[i], mx.items[j] = mx.items[j], mx.items[i]
}
func(mx *maxHeap) Len() int {
    return len(mx.items)
}