func clearStars(s string) string {
    mh := &minHeap{items: []*minHeapNode{}}
    dirtyIdx := map[int]struct{}{}
    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            dirtyIdx[i] = struct{}{}
            if mh.Len() != 0 {
                dirtyIdx[heap.Pop(mh).(*minHeapNode).idx] = struct{}{}
            }
        } else {
            heap.Push(mh, &minHeapNode{char:s[i],idx:i})
        }
    }
    out := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        _, ok := dirtyIdx[i]
        if !ok {
            out.WriteByte(s[i])
        }
    }
    return out.String()
}



type minHeapNode struct {
    char byte
    idx int
}
type minHeap struct {
	items []*minHeapNode
}

func (m *minHeap) Less(i, j int) bool {
	if m.items[i].char == m.items[j].char {
        return m.items[i].idx > m.items[j].idx
    }
    return m.items[i].char < m.items[j].char
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*minHeapNode))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
