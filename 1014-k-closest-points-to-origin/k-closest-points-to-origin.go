func kClosest(points [][]int, k int) [][]int {
    mx := &maxHeap{items: []*distNode{}}
    for i := 0; i < len(points); i++ {
        heap.Push(mx, &distNode{
            dist: calcDist(points[i][0], points[i][1]),
            point: points[i],
        })
        if mx.Len() > k {
            heap.Pop(mx)
        }
    }
    out := [][]int{}
    for i := 0; i < k; i++ {
        out = append(out, heap.Pop(mx).(*distNode).point)
    }
    return out

}

func calcDist(x , y int) float64 {
    return math.Sqrt(float64(x*x)+float64(y*y))
}

type distNode struct {
    dist float64
    point []int
}

type maxHeap struct {
	items []*distNode
}

func (m *maxHeap) Less(i, j int) bool {
	return m.items[i].dist > m.items[j].dist
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*distNode))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
