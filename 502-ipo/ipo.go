func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    merged := [][]int{}
    for i := 0; i < len(profits); i++ {
        merged = append(merged, []int{capital[i], profits[i]})
    }
    sort.Slice(merged, func(i, j int)bool{
        return merged[i][0] < merged[j][0]
    })
    i := 0
    total := w
    mx := &maxHeap{items: []int{}}
    for k > 0 {
        for i < len(merged) && w >= merged[i][0] {
            heap.Push(mx, merged[i][1])
            i++
        }
        if mx.Len() == 0 {break}
        top := heap.Pop(mx).(int)
        w += top
        total += top
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
