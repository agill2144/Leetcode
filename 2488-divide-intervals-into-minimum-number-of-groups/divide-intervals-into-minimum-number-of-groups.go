func minGroups(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    mn := &minHeap{items: []int{}}
    for i := 0; i < len(intervals); i++ {
        start, end := intervals[i][0], intervals[i][1]
        if mn.Len() != 0 && start > mn.items[0] {
            // push to existing group
            heap.Pop(mn)
            heap.Push(mn, end)
        } else {
            // new group
            heap.Push(mn, end)
        }
    }
    return mn.Len()
}

type minHeap struct {
	items []int
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i] < m.items[j]
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(int))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
