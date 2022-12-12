func lastStoneWeight(stones []int) int {
    mx := &maxHeap{items: []int{}}
    
    for i := 0; i < len(stones); i++ {
        heap.Push(mx, stones[i])
    }
    
    for mx.Len() > 1 {
        top := heap.Pop(mx).(int)
        second := heap.Pop(mx).(int)
        if top == second {
            continue
        } else {
            heap.Push(mx, top-second)
        }
    }
    
    if mx.Len() == 0 {return 0}
    return mx.items[0]

}


type maxHeap struct {
	items []int
}
func (m *maxHeap) Less(i, j int) bool {return m.items[i] > m.items[j]}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Swap(i, j int) {m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}