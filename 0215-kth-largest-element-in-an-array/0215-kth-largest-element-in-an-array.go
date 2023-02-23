func findKthLargest(nums []int, k int) int {
    mh := &minHeap{items: []int{}}
    for i := 0; i < len(nums); i++ {
        heap.Push(mh, nums[i])
        if mh.Len() > k {
            heap.Pop(mh)
        }
    }
    return mh.items[0]
}

type minHeap struct{items []int}
func(m *minHeap) Len() int {return len(m.items)}
func(m *minHeap) Less(i,j int) bool {return m.items[i] < m.items[j]}
func(m *minHeap) Swap(i,j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
func(m *minHeap) Push(x interface{}) { m.items = append(m.items, x.(int))}
func(m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}