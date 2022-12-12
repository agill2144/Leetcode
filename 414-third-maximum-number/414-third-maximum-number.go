func thirdMax(nums []int) int {
    mn := &minHeap{items: []int{}}
    
    m := map[int]struct{}{}
    for _, ele := range nums {
        m[ele] = struct{}{}
    }
    
    for ele, _ := range m {
        heap.Push(mn, ele)
        if len(mn.items) > 3 {
            heap.Pop(mn)
        }
        
    }
    if len(mn.items) >= 3 {
        return mn.items[0]
    } else {
        out := -1
        for len(mn.items) != 0 {
            out = heap.Pop(mn).(int)
        }
        return out
    }
}


type minHeap struct {
    items []int
}
func (m *minHeap) Len() int {
    return len(m.items)
}
func (m *minHeap) Swap(i, j int) {
    m.items[i],m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Less(i, j int) bool {
    return m.items[i] < m.items[j]
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
func (m *minHeap) Push(x interface{}) {
    m.items = append(m.items, x.(int))
}