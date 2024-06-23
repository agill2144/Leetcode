func longestSubarray(nums []int, limit int) int {
    ans := 0 
    mn := &minHeap{items: [][]int{}}
    mx := &maxHeap{items: [][]int{}}
    left := 0
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        heap.Push(mx, []int{i, curr})
        heap.Push(mn, []int{i, curr})
        diff := mx.items[0][1] - mn.items[0][1]
        size := i-left+1
        if diff <= limit {
            ans = max(ans, size)
        } else {
            left = min(mn.items[0][0], mx.items[0][0])+1
            for mn.Len() != 0 && (mn.items[0][0] < left ) {
                heap.Pop(mn)
            }
            for mx.Len() != 0 && mx.items[0][0] < left {
                heap.Pop(mx)
            }
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}

type maxHeap struct {items [][]int}
func (m *maxHeap) Less(i, j int) bool {return m.items[i][1] > m.items[j][1]}
func (m *maxHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *maxHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
type minHeap struct {items [][]int}
func (m *minHeap) Less(i, j int) bool {return m.items[i][1] < m.items[j][1]}
func (m *minHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Push(x interface{}) {m.items = append(m.items, x.([]int))}
func (m *minHeap) Pop() interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}
