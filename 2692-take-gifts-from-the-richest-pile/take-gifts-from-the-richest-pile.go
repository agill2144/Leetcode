func pickGifts(gifts []int, k int) int64 {
    mx := &maxHeap{items:[]int{}}
    var total int64
    for i := 0; i < len(gifts); i++ {
        heap.Push(mx, gifts[i])
        total += int64(gifts[i])
    }
    for mx.Len() != 0 && k != 0 {
        top := heap.Pop(mx).(int)
        sqrt := int(math.Sqrt(float64(top)))
        total -= int64(top-sqrt)
        heap.Push(mx, sqrt)
        k--
    }
    return total
}


type maxHeap struct {
	items []int
}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i] > m.items[j]
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(int))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
