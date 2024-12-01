type listNode struct {
    key int
    val int
    next *listNode
    prev *listNode
}

type LRUCache struct {
    head *listNode
    tail *listNode
    m map[int]*listNode
    capacity int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        m: map[int]*listNode{},
        capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.m[key]
    if !ok { return -1 }
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
        toBeDeleted := this.tail
        this.removeNode(toBeDeleted)
        delete(this.m, toBeDeleted.key)
    }
    newNode := &listNode{key:key, val:value}
    this.addToHead(newNode)
    this.m[key] = newNode
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

    // node is a head node
    if node == this.head {
        this.head = next
        if next != nil {next.prev = nil}
        return
    }

    // node is a tail node
    if node == this.tail {
        this.tail = prev
        if prev != nil {prev.next = nil}
        return
    }

    // node is somewhere in the middle
    next.prev = prev
    prev.next = next
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */