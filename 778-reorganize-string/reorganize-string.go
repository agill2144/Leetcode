func reorganizeString(s string) string {
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
    }
    mx := &maxHeap{items:[]item{}}
    for k, v := range freq {
        heap.Push(mx, item{k,v})
    }
    st := []byte{}
    for mx.Len() != 0 {
        top := heap.Pop(mx).(item)
        // can use top char
        if len(st) == 0 || st[len(st)-1] != top.char {
            st = append(st, top.char)
            top.count--
            if top.count > 0 {heap.Push(mx, top)}
            continue
        }

        // at this point, top of st == top of heap
        // we need a diff character
        // if heap is empty, we are screwed
        if mx.Len() == 0 {return ""}
        top2 := heap.Pop(mx).(item)
        heap.Push(mx, top)
        st = append(st, top2.char)
        top2.count--
        if top2.count > 0 {heap.Push(mx, top2)}
        
    }
    if len(st) != len(s) {return ""}
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {out.WriteByte(st[i])}
    return out.String()
}


type item struct {
    char byte
    count int
}
type maxHeap struct {
	items []item
}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i].count > m.items[j].count
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(item))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
