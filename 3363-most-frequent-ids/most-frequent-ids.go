func mostFrequentIDs(nums []int, freq []int) []int64 {
    out := make([]int64, len(nums))
    pq := &maxHeap{items: [][]int{}}
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        id := nums[i]
        f := freq[i]
        freqMap[id] += f
        heap.Push(pq, []int{id, freqMap[id]})
        
        for true {
            topId, topFreq := pq.items[0][0], pq.items[0][1]
            correctFreq := freqMap[topId]
            if topFreq != correctFreq {
                heap.Pop(pq)
                heap.Push(pq, []int{topId, correctFreq})
            } else {break}
        }
        
        out[i] = int64(pq.items[0][1])
    }
    return out
}

type maxHeap struct {
    items [][]int // [id, freq]
}

func (m *maxHeap) Less(i,j int) bool {
    return m.items[i][1] > m.items[j][1]
}
func (m *maxHeap) Swap(i,j int) {
    m.items[i],m.items[j] = m.items[j],m.items[i]
}
func (m *maxHeap) Len() int {
    return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
    m.items = append(m.items, x.([]int))
}
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}