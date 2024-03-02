type LRUCache struct {
    capacity int
    dll *dll
    data map[int]*listNode
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        data: map[int]*listNode{},
        dll: &dll{},
    }
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.data[key]
    if !ok {return -1}
    this.dll.moveNodeToHead(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    
    node, ok := this.data[key]
    if ok {
        node.val = value
        this.dll.moveNodeToHead(node)
        return
    }
    if len(this.data) == this.capacity {
        lruNode := this.dll.deleteFromTail()
        delete(this.data, lruNode.key)
    }
    this.data[key] = this.dll.addToHead(key, value)
}


type listNode struct {
    key int
    val int
    prev *listNode
    next *listNode
}

type dll struct {
    head *listNode // newely added or recently updated
    tail *listNode // least recently used
}

func (d *dll) addToHead(key, val int) *listNode {
    newNode := &listNode{val: val, key: key}
    if d.head == nil {
        d.head = newNode
        d.tail = newNode
        return newNode
    }
    currHead := d.head
    currHead.prev = newNode
    newNode.next = currHead
    d.head = newNode
    return newNode
}

func (d *dll) deleteFromTail() *listNode {
    currTail := d.tail
    if d.head == d.tail {
        d.head = nil
        d.tail = nil
        return currTail
    }

    newTail := currTail.prev
    newTail.next = nil
    currTail.prev = nil
    d.tail = newTail
    return currTail
}

func (d *dll) moveNodeToHead(node *listNode) {
    if node == d.head {return}
    prev := node.prev
    next := node.next
        node.prev = nil
        node.next = nil

    if node == d.tail {
        prev.next = nil
        d.tail = prev    
    } else {
        prev.next = next
        next.prev = prev
    }

    node.next = d.head
    d.head.prev = node
    d.head = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */