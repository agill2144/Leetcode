type LRUCache struct {
    dll *dll
    capacity int
    cache map[int]*listNode
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        cache: map[int]*listNode{},
        capacity: capacity,
        dll: &dll{},
    }
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.cache[key]
    if ok {
        this.dll.removeNode(node)
        this.dll.append(node)
        return node.val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.cache[key]
    if ok {
        node.val = value
        this.dll.removeNode(node)
        this.dll.append(node)
        return
    }
    if len(this.cache) == this.capacity {
        toBeDeleted := this.dll.head
        this.dll.removeNode(toBeDeleted)
        delete(this.cache, toBeDeleted.key)
    }
    newNode := &listNode{key:key,val:value}
    this.dll.append(newNode)
    this.cache[key] = newNode
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
    head *listNode // old / stale side of data
    tail *listNode // fresh / latest side of data
}

func (d *dll) append(node *listNode) {
    if d.head == nil {
        d.head = node
        d.tail = node
        return
    }
    d.tail.next = node
    node.prev = d.tail
    d.tail = node
}

func (d *dll) removeNode(node *listNode) {
    next := node.next
    prev := node.prev
    node.next = nil
    node.prev = nil
    if next != nil {next.prev = prev}
    if prev != nil {prev.next = next}
    if d.head == node {
        d.head = next
    }
    if d.tail == node {
        d.tail = prev
    }
}
