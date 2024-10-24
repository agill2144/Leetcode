type stNode struct {
    val int
}

type MaxStack struct {
    st []*stNode
    mx *maxHeap
    removed map[*stNode]bool
}


func Constructor() MaxStack {
    return MaxStack{
        st: []*stNode{},
        mx: &maxHeap{items: []*stNode{}},
        removed: map[*stNode]bool{},
    }
}


func (this *MaxStack) Push(x int)  {
    node := &stNode{x}
    heap.Push(this.mx, node)
    this.st = append(this.st, node)
}


func (this *MaxStack) Pop() int {
    this.adjustSt()
    out := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    this.removed[out] = true
    return out.val
}


func (this *MaxStack) Top() int {
    this.adjustSt()
    return this.st[len(this.st)-1].val
}


func (this *MaxStack) PeekMax() int {
    this.adjustHeap()
    return this.mx.items[0].val
}


func (this *MaxStack) PopMax() int {
    this.adjustHeap()
    topMax := heap.Pop(this.mx).(*stNode)
    this.removed[topMax] = true
    return topMax.val
}

func (this *MaxStack) adjustSt() {
    removed := this.removed[this.st[len(this.st)-1]]
    for len(this.st) != 0 && removed {
        this.st = this.st[:len(this.st)-1]
        removed = this.removed[this.st[len(this.st)-1]]
    }
}

func (this *MaxStack) adjustHeap() {
    for this.mx.Len() != 0 && this.removed[this.mx.items[0]] {
        heap.Pop(this.mx)
    }
}

type maxHeap struct {items []*stNode}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i].val >= m.items[j].val
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*stNode))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
