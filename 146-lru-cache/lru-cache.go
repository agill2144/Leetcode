type listNode struct {
    key int
    val int
    next *listNode
    prev *listNode
}

func (this *LRUCache) addToHead(node *listNode) {
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

    // when node is head node
    if node == this.head {
        this.head = next
        if next != nil {next.prev = nil}
        return
    }

    // when node is tail node
    if node == this.tail {
        prev.next = nil
        this.tail = prev
        return
    }


    // when node is some middle node
    next.prev = prev
    prev.next = next
}



type LRUCache struct {
    head *listNode
    tail *listNode
    m map[int]*listNode
    capacity int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{m: map[int]*listNode{}, capacity: capacity}
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.m[key]
    if !ok {return -1}
    this.removeNode(node)
    this.addToHead(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.m[key]
    if ok {
        node.val = value
        this.removeNode(node)
        this.addToHead(node)
        return
    }
    if len(this.m) == this.capacity {
        deletedTail := this.tail
        this.removeNode(deletedTail)
        delete(this.m, deletedTail.key)
    }
    newNode := &listNode{key:key, val:value}
    this.addToHead(newNode)
    this.m[key] = newNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */