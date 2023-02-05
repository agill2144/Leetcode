func kthSmallest(matrix [][]int, k int) int {
    mx := &maxHeap{items: []int{}}
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            heap.Push(mx, matrix[i][j])
            if len(mx.items) > k {
                heap.Pop(mx)
            }
        }
    }
    return mx.items[0]
}


type maxHeap struct {
    items []int
}

func (m *maxHeap) Less(i, j int) bool {return m.items[i] > m.items[j]}
func (m *maxHeap) Swap(i, j int) {m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Push(x interface{}) { m.items = append(m.items, x.(int))}
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}