func findClosestElements(arr []int, k int, x int) []int {
    mx := &maxHeap{items: [][]int{}}    
    for i := 0; i < len(arr); i++ {
        dist := abs(arr[i] - x)
        heap.Push(mx, []int{dist, i})
        if mx.Len() > k {heap.Pop(mx)}
    }
    out := []int{}
    for mx.Len() != 0 {
        out = append(out, arr[heap.Pop(mx).([]int)[1]])
    }
    sort.Ints(out)
    return out
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}


type maxHeap struct {
    items [][]int // [ [dist, idx] ]
}

func (m *maxHeap) Less(i, j int) bool {
    iIdx, jIdx := m.items[i][1], m.items[j][1]
    iDist, jDist := m.items[i][0], m.items[j][0]
    if iDist == jDist {
        return iIdx > jIdx
    }
    return iDist > jDist
}
func (m *maxHeap) Swap(i, j int) {
    m.items[i], m.items[j] = m.items[j], m.items[i]
}

func (m *maxHeap) Len() int {
    return len(m.items)
}

func (m *maxHeap) Pop()interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

func (m *maxHeap) Push(x interface{}) {
    m.items = append(m.items, x.([]int))
} 