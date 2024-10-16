func longestDiverseString(a int, b int, c int) string {
    mx := &maxHeap{items: []*item{}}
    if a > 0 {heap.Push(mx, &item{'a',a})}
    if b > 0 {heap.Push(mx, &item{'b',b})}
    if c > 0 {heap.Push(mx, &item{'c',c})}
    out := new(strings.Builder)
    for mx.Len() != 0 {
        top := heap.Pop(mx).(*item)
        char := top.val
        freq := top.freq
        outStr := out.String()
        // we cannot use this char if prev and prevPrev == curr char
        if len(outStr) >= 2 && 
            outStr[len(outStr)-1] == char &&
            outStr[len(outStr)-2] == char {

                if mx.Len() == 0 {break}
                t := top
                top = heap.Pop(mx).(*item)
                heap.Push(mx, t)
                char = top.val
                freq = top.freq
        }
        
        out.WriteByte(char)
        freq--
        if freq > 0 {heap.Push(mx, &item{char, freq})}
    }
    return out.String()
}


type item struct {
    val byte
    freq int
}

type maxHeap struct {
	items []*item
}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i].freq > m.items[j].freq
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*item))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
