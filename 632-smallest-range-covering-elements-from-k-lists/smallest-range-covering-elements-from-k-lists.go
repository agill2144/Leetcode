func smallestRange(nums [][]int) []int {
    mn := &minHeap{items: [][]int{}}
    maxVal := math.MinInt64
    // k = num of lists
    // n = avg size of each list

    // o(k)
    for i := 0; i < len(nums); i++ {
        heap.Push(mn, []int{i, 0, nums[i][0]})
        maxVal = max(maxVal, nums[i][0])
    }
    out := []int{0, math.MaxInt64}
    // o(min(n) * logk)
    for true {
        minEle := heap.Pop(mn).([]int)
        listIdx, innerIdx, c := minEle[0], minEle[1], minEle[2]
        rangeDiff := maxVal-c
        a, b:= out[0], out[1]
        currRangeDiff := b-a
        if rangeDiff < currRangeDiff || (rangeDiff == currRangeDiff && c < a){
            out = []int{c,maxVal}
        }
        if innerIdx+1 == len(nums[listIdx]) {break}
        innerIdx++
        newItem := []int{listIdx, innerIdx, nums[listIdx][innerIdx]}
        maxVal = max(maxVal, newItem[2])
        heap.Push(mn, newItem)
    }
    return out
}

type minHeap struct {
	items [][]int // {listIdx, innerIdx, val}
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i][2] < m.items[j][2]
}
func (m *minHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}
func (m *minHeap) Len() int {
	return len(m.items)
}
func (m *minHeap) Push(x interface{}) {
	m.items = append(m.items, x.([]int))
}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
