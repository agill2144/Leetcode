func maxAverageRatio(classes [][]int, extraStudents int) float64 {
    mx := &maxHeap{items: []*item{}}
    for i := 0; i < len(classes); i++ {
        p := float64(classes[i][0])
        c := float64(classes[i][1])
        heap.Push(mx, &item{p, c, calcGain(p, c)})
    }
    for extraStudents != 0 {
        top := heap.Pop(mx).(*item)
        p := top.pass + 1.0
        c := top.total + 1.0
        heap.Push(mx, &item{p,c, calcGain(p, c)})
        extraStudents--
    }
    var sum float64
    for mx.Len() != 0 {
        top := heap.Pop(mx).(*item)
        p := float64(top.pass)
        c := float64(top.total)
        sum = sum + (p/c)
    }
    return sum / float64(len(classes))
}

func calcGain(p, c float64) float64 {
    return ((p+1.0) / (c+1.0)) - (p/c)

}


type item struct {
    pass float64
    total float64
    gain float64
}

type maxHeap struct {
	items []*item
}

func (m *maxHeap) Less(i, j int) bool {
    return m.items[i].gain > m.items[j].gain
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*item))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
