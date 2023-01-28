type LRUCache struct {
    cache map[int]*listNode
    dll *dll
    max int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        cache: map[int]*listNode{},
        dll: &dll{},
        max: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    nodeRef, ok := this.cache[key]
    if !ok {
        return -1
    }
    this.update(nodeRef)
    return nodeRef.val
}

func (this *LRUCache) update(nodeRef *listNode) {
    // move this node to head
    // 1-2-3 move 3
    // 1-2-3 move 2
    if this.dll.head == nodeRef {
        return
    }
    this.dll.deleteNode(nodeRef)
    this.dll.addNodeToHead(nodeRef)
}


func (this *LRUCache) Put(key int, value int)  {
    nodeRef, ok := this.cache[key]
    if ok {
        nodeRef.val = value
        this.update(nodeRef)
    } else {
        if len(this.cache) == this.max {
            nodeToDelete := this.dll.tail
            this.dll.deleteNode(nodeToDelete)
            delete(this.cache, nodeToDelete.key)
        }
        newNode := &listNode{key: key, val: value}
        this.dll.addNodeToHead(newNode)
        this.cache[key] = newNode
    }
}

type listNode struct {
    key int
    val int
    next *listNode
    prev *listNode
}

type dll struct {
    head *listNode
    tail *listNode
    size int
}

func (d *dll) addNodeToHead(nodeRef *listNode) {
    if d.head == nil {
        d.head = nodeRef
        d.tail = nodeRef
        d.size = 1
        return
    }
    
    nodeRef.next = d.head
    d.head.prev = nodeRef
    d.head = nodeRef
    d.size++
}

func (d *dll) deleteNode(nodeRef *listNode) {
    if d.head == nil || nodeRef == nil {
        return
    }
    next := nodeRef.next
    prev := nodeRef.prev
    
    // 1 -- if nodeRef is the node to delete
    if next == nil && prev == nil {
        d.head = nil
        d.tail = nil
        d.size = 0
        return
    }
    
    if prev == nil {
        // this is the head node
        // 1<->2<->3 ( deleting head )     
        newHead := next
        d.head.next = nil
        newHead.prev = nil
        d.head = newHead
    } else {
        // 1-2-3 ( deleting tail )
        // 1-2-3 ( deleting node in the middle )
        prev.next = next
        if next != nil {
            next.prev = prev
        } else {
            d.tail = prev
        }
    }
    d.size--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */