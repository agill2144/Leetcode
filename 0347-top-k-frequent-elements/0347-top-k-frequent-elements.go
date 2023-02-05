func topKFrequent(nums []int, k int) []int {
    freqMap := map[int]int{}
    for _, el := range nums {
        freqMap[el]++
    }
    mn := &minHeap{items: []*item{}}
    for num, count := range freqMap {
        heap.Push(mn, &item{val: num ,freq: count})
        if len(mn.items) > k {
            heap.Pop(mn)
        }
    }
    out := []int{}
    for len(mn.items) != 0 {
        out = append(out, heap.Pop(mn).(*item).val)
    }
    return out
}

type minHeap struct {
    items []*item
}
type item struct {
    val int
    freq int
}

func (m *minHeap) Less(i,j int) bool {
    return m.items[i].freq < m.items[j].freq
}
func (m *minHeap) Swap(i,j int) {
    m.items[i],m.items[j] = m.items[j],m.items[i]
}
func (m *minHeap) Len() int {
    return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
    m.items = append(m.items, x.(*item))
}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}
