type AllOne struct {
    freqLL map[int]*ll
    nodes map[string]*listNode
    maxFreq int
}


func Constructor() AllOne {
    return AllOne{
        freqLL: map[int]*ll{},
        nodes: map[string]*listNode{},
        maxFreq: 0,
    }
}


func (this *AllOne) Inc(key string)  {
    node, ok := this.nodes[key]
    if ok {
        ll := this.freqLL[node.val]
        ll.remove(node)
        node.val++
        if this.freqLL[node.val] == nil {
            this.freqLL[node.val] = newLL()
        }
        this.freqLL[node.val].append(node)
        this.maxFreq = max(this.maxFreq, node.val)
        return
    }
    node = &listNode{val: 1, key: key}
    if this.freqLL[node.val] == nil {
        this.freqLL[node.val] = newLL()
    }
    this.freqLL[node.val].append(node)
    this.nodes[key] = node
    this.maxFreq = max(this.maxFreq, node.val)
}


func (this *AllOne) Dec(key string)  {
    node, ok := this.nodes[key]
    if ok {
        ll := this.freqLL[node.val]
        ll.remove(node)
        node.val--
        if node.val > 0 {
            this.freqLL[node.val].append(node)
        } else {
            delete(this.nodes, key)
        }
        if this.maxFreq == node.val && ll.head == nil {this.maxFreq--}
        if this.maxFreq == -1 {this.maxFreq = 0}
    }
}


func (this *AllOne) GetMaxKey() string {
    for i := this.maxFreq; i >= 0; i-- {
        if this.freqLL[i] != nil &&
            this.freqLL[i].head != nil {return this.freqLL[i].head.key}
    }
    return ""    
}


func (this *AllOne) GetMinKey() string {
    for i := 1; i <= this.maxFreq; i++ {
        if this.freqLL[i] != nil &&
            this.freqLL[i].head != nil {return this.freqLL[i].head.key}
    }
    return ""
}


type listNode struct {
    key string
    val int
    prev *listNode
    next *listNode
}

type ll struct {
    head *listNode
    tail *listNode
}

func newLL() *ll {
    return new(ll)
}

func (l *ll) append(node *listNode) {
    if l.head == nil {
        l.head = node
        l.tail = node
        return
    }
    l.tail.next = node
    node.prev = l.tail
    l.tail = node
}


func (l *ll) remove(node *listNode) {
    next := node.next
    prev := node.prev
    node.next = nil
    node.prev = nil
    if node == l.head {
        newHead := next
        l.head = newHead
        if next != nil {next.prev = nil}
        return
    }
    if l.tail == node {
        l.tail = prev
        if prev != nil {prev.next = nil}
        return
    }

    prev.next = next
    if next != nil {next.prev = prev}
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */