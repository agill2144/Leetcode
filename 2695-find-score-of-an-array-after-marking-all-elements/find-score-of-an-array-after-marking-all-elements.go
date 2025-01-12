func findScore(nums []int) int64 {
    mn := &minHeap{items: [][]int{}}
    for i := 0; i < len(nums); i++ {
        heap.Push(mn, []int{nums[i], i})
    }
    var score int64
    usedIdxs := map[int]bool{}
    for mn.Len() != 0 {
        top := heap.Pop(mn).([]int)
        val,idx := top[0], top[1]
        if usedIdxs[idx] {continue}
        score += int64(val)
        usedIdxs[idx-1] = true
        usedIdxs[idx+1] = true
    }
    return score
}

type minHeap struct {
	items [][]int // [val, idx]
}

func (m *minHeap) Less(i, j int) bool {
	if m.items[i][0] == m.items[j][0] {
        return m.items[i][1] < m.items[j][1]
    }
	return m.items[i][0] < m.items[j][0]
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.([]int))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
