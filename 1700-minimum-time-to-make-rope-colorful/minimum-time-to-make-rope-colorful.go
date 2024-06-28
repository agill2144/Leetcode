func minCost(colors string, neededTime []int) int {
    totalTime := 0
    n := len(colors)
    pq := &minHeap{items: []int{}}
    i := 0
    for i < n-1 {
        for i < n-1 && colors[i] == colors[i+1] {
            heap.Push(pq, neededTime[i])
            i++
        }
        heap.Push(pq, neededTime[i])        
        
        for pq.Len() > 1 {
            minTime := heap.Pop(pq).(int)
            totalTime += minTime            
        }
        heap.Pop(pq)
        i++   
    }
    return totalTime
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
