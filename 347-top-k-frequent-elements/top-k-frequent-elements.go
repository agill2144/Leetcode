func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}

    bucket := make([][]int, len(nums)+1)
    for num, count := range freq {
        if bucket[count] == nil {bucket[count] = []int{}}
        bucket[count] = append(bucket[count], num)
    }
    out := []int{}
    for i := len(bucket)-1; i >= 0 && len(out) != k; i-- {
        if bucket[i] == nil {continue}
        out = append(out, bucket[i]...)
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
