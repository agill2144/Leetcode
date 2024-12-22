type listNode struct {
    key int
    val int
    next *listNode
    prev *listNode
}

type LRUCache struct {
    data map[int]*listNode
    head *listNode // most recent node , append new nodes here
    tail *listNode // least recently used node, discard from tail
    cap int    
}

func (this *LRUCache) addNode(node *listNode) {
    if this.head == nil {
        this.head = node
        this.tail = node
        return
    }
    node.next = this.head
    this.head.prev = node
    this.head = node
}

func (this *LRUCache) removeNode(node *listNode) {
    next := node.next
    prev := node.prev
    node.next = nil
    node.prev = nil
    // prev -> next ( if prev not nil )
    // next <- prev ( if next not nil )
    if prev != nil {prev.next = next}
    if next != nil {next.prev = prev}

    // update head ptr
    if node == this.head {
        this.head = next
        return
    }
    
    // update tail ptr
    if node == this.tail {
        this.tail = prev
        return
    }
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        data: map[int]*listNode{},
        cap: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.data[key]
    if !ok {return -1}
    this.removeNode(node)
    this.addNode(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.data[key]
    if ok {
        node.val = value
        this.removeNode(node)
        this.addNode(node)
        return
    }
    if len(this.data) == this.cap {
        toBeDeleted := this.tail
        this.removeNode(toBeDeleted)
        delete(this.data, toBeDeleted.key)
    }

    newNode := &listNode{key:key, val:value}
    this.addNode(newNode)
    this.data[key] = newNode
    
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */