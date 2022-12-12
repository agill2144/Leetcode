

type item struct {
	val int
	numTimes int
}
type minHeap struct {
	items []*item
}
func (m *minHeap) Less(i, j int) bool {return m.items[i].numTimes < m.items[j].numTimes}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*item))
}

func topKFrequent(nums []int, k int) []int {
	if nums == nil {
		return nums
	}
	m := map[int]int{}
	for _, e := range nums {
		m[e]++
	}
	mn := &minHeap{items: []*item{}}
	for num, times := range m {
		i := &item{val: num, numTimes: times}
		if len(mn.items) != k {
			heap.Push(mn, i)
			continue
		}
		peek := mn.items[0]
		if times > peek.numTimes {
			heap.Pop(mn)
			heap.Push(mn, i)
		}
	}
	out := []int{}
	for len(mn.items) != 0 {
		out = append(out, heap.Pop(mn).(*item).val)
	}
	return out
}
