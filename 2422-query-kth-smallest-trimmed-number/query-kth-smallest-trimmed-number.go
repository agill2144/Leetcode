func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
    n := len(nums)
    q := len(queries)
    out := []int{}
    for i := 0; i < q; i++ {

        query := queries[i]
        keepLast := query[1]
        kth := query[0]
        mx := &maxHeap{items: []*mxNode{}}

        for j := 0; j < n; j++ {
            num := nums[j]
            subStr := num[len(num)-keepLast:]
            // cant do this, apparently value goes out of bound 
            // panic: strconv.Atoi: parsing "9288488870527604910029": value out of range
            // subStrInt, _ := strconv.Atoi(subStr)
            heap.Push(mx, &mxNode{val: subStr, idx: j})
            if mx.Len() > kth {
                heap.Pop(mx)
            }
        }
        out = append(out, mx.items[0].idx)
    }
    return out
}

type mxNode struct {
    val string
    idx int
}
type maxHeap struct {
	items []*mxNode
}

func (m *maxHeap) Less(i, j int) bool {
	if m.items[i].val == m.items[j].val {
        return m.items[i].idx > m.items[j].idx
    }
    return m.items[i].val > m.items[j].val
}
func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *maxHeap) Len() int {
	return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
	m.items = append(m.items, x.(*mxNode))
}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
