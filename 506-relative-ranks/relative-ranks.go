func findRelativeRanks(score []int) []string {
    n := len(score)
    mx := &maxHeap{items: [][]int{}}
    for i := 0; i < n; i++ {
        heap.Push(mx, []int{score[i],i})
    }
    nr := 1
    out := make([]string, n)
    for mx.Len() != 0 {
        popped := heap.Pop(mx).([]int)
        idx := popped[1]
        if nr == 1 {
            out[idx] = "Gold Medal"
        } else if nr == 2 {
            out[idx] = "Silver Medal"
        } else if nr == 3 {
            out[idx] = "Bronze Medal"
        } else {
            out[idx] = fmt.Sprintf("%v", nr)
        }
        nr++
    }
    return out
}


type maxHeap struct {
	items [][]int // val, idx
}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i][0] > m.items[j][0]
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.([]int))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
