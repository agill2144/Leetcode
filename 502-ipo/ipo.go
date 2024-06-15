func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
    merged := [][]int{}
    for i := 0; i < len(profits); i++ {
        merged = append(merged, []int{capital[i], profits[i]})
    }
    sort.Slice(merged, func(i, j int)bool{
        return merged[i][0] < merged[j][0]
    })
    i := 0
    total := w
    mx := &maxHeap{items: []int{}}
    for k > 0 {

        //1. find out which projects we can take 
        // - projects that need <= our current capital; we can take it
        // find all of them first
        for i < len(merged) && w >= merged[i][0] {
            heap.Push(mx, merged[i][1])
            i++
        }
        if mx.Len() == 0 {break}

        //2. out of all of the ones we can take,
        // we want to process the project that maximizes our profit
        // therefore we need heap to keep track of max profit
        // take the top, reduce k since we have taken a project now
        // KEEP IN MIND, remaining projects that we did not pop, are still in heap
        // they will be compared again until we have taken all k projects
        top := heap.Pop(mx).(int)
        w += top
        total += top
        k--
    }
    return total
    
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
