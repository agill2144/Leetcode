type LFUCache struct {
    keyToNode map[int]*listNode
    useCountToDll map[int]*dll
    capacity int
    min int
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        keyToNode: map[int]*listNode{},
        useCountToDll: map[int]*dll{},
        capacity: capacity,
        min: 0,
    }
}


func (this *LFUCache) Get(key int) int {
    nodeRef, ok := this.keyToNode[key]
    if !ok {
        return -1
    }
    this.update(nodeRef)
    return nodeRef.val
}

func (this *LFUCache) update(nodeRef *listNode) {
    currUseCount := nodeRef.count
    currDll := this.useCountToDll[currUseCount]
    currDll.removeNode(nodeRef)
    if currUseCount == this.min && currDll.size == 0 {
        this.min++
    }
    
    nodeRef.count++
    targetDll, exists := this.useCountToDll[nodeRef.count]
    if !exists {
        this.useCountToDll[nodeRef.count] = new(dll)
        this.useCountToDll[nodeRef.count].addToHead(nodeRef)
    } else {
        targetDll.addToHead(nodeRef)
    }
}



func (this *LFUCache) Put(key int, value int)  {
    if this.capacity == 0 {return}
    nodeRef, ok := this.keyToNode[key]
    if ok {
        nodeRef.val = value
        this.update(nodeRef)
    } else {
        if len(this.keyToNode) == this.capacity {
            targetDll := this.useCountToDll[this.min]
            toDelete := targetDll.tail
            targetDll.removeNode(toDelete)
            delete(this.keyToNode, toDelete.key)
        }
        this.min = 1
        newNode := &listNode{key: key, val: value, count: 1}
        this.keyToNode[key] = newNode
        targetDll, exists := this.useCountToDll[1]
        if !exists {
            this.useCountToDll[1] = new(dll)
            this.useCountToDll[1].addToHead(newNode)
        } else {
            targetDll.addToHead(newNode)
        }
    }
}


type listNode struct {
    key int
    val int
    next *listNode
    prev *listNode
    count int
}

type dll struct {
    head *listNode
    tail *listNode
    size int
}

func (d *dll) addToHead(n *listNode) {
    if d.head == nil {
        d.head = n
        d.tail = n
        d.size = 1
        return
    }
    newHead := n
    newHead.next = d.head
    d.head.prev = newHead
    d.head = newHead
    d.size++
}


func (d *dll) removeNode(n *listNode) {
    if d.head == nil || n == nil {
        return
    }
    next := n.next
    prev := n.prev
    
    if next == nil && prev == nil {
        d.head = nil
        d.tail = nil
        d.size = 0
        return
    }
    
    if prev != nil {
        prev.next = next
        if next != nil {
            next.prev = prev
        } else {
            d.tail = prev
        }
    } else {
        newHead := next
        newHead.prev = nil
        d.head.next = nil
        d.head = newHead
    }
    d.size--
    n.next = nil
    n.prev = nil
}
/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */