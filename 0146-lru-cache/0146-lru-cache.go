type LRUCache struct {
    doublyLL *dll
    k int
    index map[int]*listNode
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        k: capacity,
        doublyLL: &dll{},
        index: map[int]*listNode{},
    }
}


func (this *LRUCache) Get(key int) int {
    node, found := this.index[key]
    if !found {return -1}
    
    this.doublyLL.moveNodeToTail(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, found := this.index[key]
    if found {
        node.val = value
        this.doublyLL.moveNodeToTail(node)
        return
    }
    if len(this.index) == this.k {
        deletedNode := this.doublyLL.deleteHead()
        delete(this.index, deletedNode.key)
    }
    addedNode := this.doublyLL.appendNode(key, value)
    this.index[key] = addedNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


type listNode struct {
    key int
    val int
    next *listNode
    prev *listNode
    
}

type dll struct {
    head *listNode
    tail *listNode
}


func (this *dll) appendNode(k, v int) *listNode {
    newTail := &listNode{key:k, val: v}
    if this.head == nil {
        this.head = newTail
        this.tail = newTail
        return newTail
    }
    this.tail.next = newTail
    newTail.prev = this.tail
    this.tail = newTail
    return newTail
}

func (this *dll) deleteHead() *listNode {
    out := this.head
    if this.head == nil || this.head.next == nil {
        this.head = nil
        this.tail = nil
        return out
    }

    newHead := this.head.next
    newHead.prev = nil
    this.head.next = nil
    this.head = newHead
    return out
}


func (this *dll) moveNodeToTail(node *listNode) {
    if this.head == nil {panic("head is nil, where did you get this node from...")}
    if this.tail == node {
        return
    }
    prev := node.prev
    next := node.next
    node.next = nil
    node.prev = nil
    
    //   H
    // 1 2<->3
    if prev == nil {
        node.next = nil
        next.prev = nil
        this.head = next

    // 1<->3 2
    } else {
        prev.next = next
        next.prev = prev
    }
    
    this.tail.next = node
    node.prev = this.tail
    this.tail = node
}