func findClosestElements(arr []int, k int, x int) []int {
    mx := &maxHeap{items: []*node{}}    
    for i := 0; i < len(arr); i++ {
        node := &node{
            val: arr[i],
            dist: abs(arr[i]-x),
        }
        heap.Push(mx, node)
        if mx.Len() > k {heap.Pop(mx)}
    }
    out := []int{}
    for mx.Len() > 0 {
        out = append(out, heap.Pop(mx).(*node).val)
    }
    sort.Ints(out)
    return out
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}

type node struct {
    val int
    dist int
}

type maxHeap struct {
    items []*node
}

func (m *maxHeap) Less(i, j int)bool{
    if m.items[i].dist == m.items[j].dist {
        return m.items[i].val > m.items[j].val
    }
    return m.items[i].dist > m.items[j].dist
}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Swap(i, j int){
    m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Push(x interface{}) {
    m.items = append(m.items, x.(*node))
}
func (m *maxHeap) Pop()interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}