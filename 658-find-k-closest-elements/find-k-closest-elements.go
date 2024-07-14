/*
    approach: heap
    - find k closest
    - heap vibes
    - closest = smallest = max heap
        - we will keep only k elements
        - sort by dist
        - if two distances are same, sort by value

    time = o(nlogk) + o(k) + o(klogk)
    space = o(k)


*/
func findClosestElements(arr []int, k int, x int) []int {
    mx := &maxHeap{items: []*heapNode{}}
    // o(n * logk)
    for i := 0; i < len(arr); i++ {
        val := arr[i]
        dist := abs(val-x)
        heap.Push(mx, &heapNode{dist:dist, val:val})
        if mx.Len() > k {
            heap.Pop(mx)
        }
    }
    out := []int{}
    // o(k)
    for mx.Len() != 0 {
        out = append(out, heap.Pop(mx).(*heapNode).val)
    }
    // o(klogk)
    sort.Ints(out)
    return out
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}

type heapNode struct {
    dist int
    val int
}
type maxHeap struct {
	items []*heapNode
}

func (m *maxHeap) Less(i, j int) bool {
	if m.items[i].dist == m.items[j].dist {
        return m.items[i].val > m.items[j].val
    }
    return m.items[i].dist > m.items[j].dist
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*heapNode))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
