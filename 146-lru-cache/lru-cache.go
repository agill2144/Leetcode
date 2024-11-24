type LRUCache struct {
    capacity int
    head *listNode
    tail *listNode
    m map[int]*listNode
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        m: map[int]*listNode{},
    }    
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.m[key]
    if !ok {return -1}
    this.removeNode(node)
    this.addToHead(node)
    return node.value    
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.m[key]
    if ok {
        node.value = value
        this.removeNode(node)
        this.addToHead(node)
        return
    }
    if len(this.m) == this.capacity {
        toBeDeleted := this.m[this.tail.key]
        this.removeNode(toBeDeleted)
        delete(this.m, toBeDeleted.key)
    }
    newHead := &listNode{key:key,value:value}
    this.addToHead(newHead)
    this.m[key] = newHead
}

func (this *LRUCache) printLL() {
    msg := "Head->"
    curr := this.head
    for curr != nil {
        msg += fmt.Sprintf("%v", curr.value)
        if curr.next != nil {msg += "-"}
        curr = curr.next
    }
    msg += "<-Tail"
    fmt.Println(msg)
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
    if node == nil {return}
    prev := node.prev
    next := node.next
    node.prev = nil
    node.next = nil

    // when this is a head node
    if this.head == node {
        this.head = next
        if next != nil {next.prev = nil}
        return        
    }

    // when this is a tail node
    if node == this.tail {
        prev.next = nil
        this.tail = prev
        return
    }

    // when this is a node in the middle
    prev.next = next
    next.prev = prev
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type listNode struct {
    key int
    value int
    next *listNode
    prev *listNode
}
