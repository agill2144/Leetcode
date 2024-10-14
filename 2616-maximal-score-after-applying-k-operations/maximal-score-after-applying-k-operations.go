func maxKelements(nums []int, k int) int64 {
    mx := &maxHeap{items: []int{}}
    for i := 0; i < len(nums); i++ {
        heap.Push(mx, nums[i])
    }
    var score int64
    for k != 0 && mx.Len() != 0 {
        top := heap.Pop(mx).(int)
        if top <= 0 {break}
        score += int64(top)
        newItem := int(math.Ceil(float64(top) / 3.0))
        if newItem > 0 {heap.Push(mx, newItem)}
        k--
    }
    return score
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
