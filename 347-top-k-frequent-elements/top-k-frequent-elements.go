func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}

    mn := &minHeap{items: [][]int{}}
    for key, val := range freq {
        heap.Push(mn, []int{key, val})
        if mn.Len() > k {heap.Pop(mn)}
    }
    out := []int{}
    for i := 0; i < mn.Len(); i++ {
        out = append(out, mn.items[i][0])
    }
    return out
}

type minHeap struct {
	items [][]int // [ [val, freq] ]
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i][1] < m.items[j][1]
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
