func topKFrequent(nums []int, k int) []int {
    n := len(nums)
    freq := map[int]int{}
    for i := 0; i < n; i++ {
        freq[nums[i]]++
    }
    bucket := make([][]int, n+1)
    for k, v := range freq {
        if bucket[v] == nil {bucket[v] = []int{}}
        bucket[v] = append(bucket[v],k)
    }
    out := []int{}
    for i := n; i >= 0 && len(out) != k; i-- {
        if bucket[i] != nil {
            out = append(out, bucket[i]...)
        }
    }
    return out
}
// tc = o(n) + o(n*logk) + o(k)
// sc = o(n) + o(k)
// func topKFrequent(nums []int, k int) []int {
//     freq := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         freq[nums[i]]++
//     }
//     mn := &minHeap{items: [][]int{}}
//     for key,v := range freq {
//         heap.Push(mn, []int{key,v})
//         if mn.Len() > k {
//             heap.Pop(mn)
//         }
//     }
//     out := []int{}
//     for mn.Len() != 0 {
//         out = append(out, heap.Pop(mn).([]int)[0])
//     }
//     return out
// }



// type minHeap struct {
// 	items [][]int // <val, freq>
// }

// func (m *minHeap) Less(i, j int) bool {
// 	return m.items[i][1] < m.items[j][1]
// }
// func (m *minHeap) Swap(i, j int) {
// 	m.items[i], m.items[j] = m.items[j], m.items[i]
// }
// func (m *minHeap) Len() int {
// 	return len(m.items)
// }
// func (m *minHeap) Push(x interface{}) {
// 	m.items = append(m.items, x.([]int))
// }
// func (m *minHeap) Pop() interface{} {
// 	out := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return out
// }
