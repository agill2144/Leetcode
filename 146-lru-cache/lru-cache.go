type LRUCache struct {
    dll *dll
    data map[int]*listNode
    cap int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        dll: &dll{},
        data: map[int]*listNode{},
        cap: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.data[key]
    if !ok {return -1}
    this.dll.moveNodeToEnd(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.data[key]
    if ok {
        node.val = value
        this.dll.moveNodeToEnd(node)
        return
    }
    if len(this.data) == this.cap {
        deletedNode := this.dll.removeHead()
        delete(this.data, deletedNode.key)
    }
    newNode := this.dll.appendToEnd(key, value)
    this.data[key] = newNode
}


type dll struct {
    head *listNode
    tail *listNode
}

type listNode struct {
    key int
    val int
    prev *listNode
    next *listNode
}

func (d *dll) appendToEnd(key, val int) *listNode {
    newNode := &listNode{key: key, val: val}
    if d.head == nil {
        d.head = newNode
        d.tail = newNode
        return newNode
    }
    d.tail.next = newNode
    newNode.prev = d.tail
    d.tail = newNode
    return newNode
}

func (d *dll) removeHead() *listNode {
    if d.head == nil {return nil}
    out := d.head
    if d.head == d.tail {
        d.head = nil
        d.tail = nil
        return out
    }
    newHead := d.head.next
    newHead.prev = nil
    d.head.next = nil 
    d.head = newHead
    return out
}

func (d *dll) moveNodeToEnd(node *listNode) {
    if node == d.tail {return}    
    if node == d.head {
        newHead := d.head.next
        newHead.prev = nil
        node.next = nil
        d.head = newHead
    } else {
        prev := node.prev
        next := node.next
        prev.next = next
        next.prev = prev
        node.next = nil
        node.prev = nil
    }
    node.prev = d.tail
    d.tail.next = node
    d.tail = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */