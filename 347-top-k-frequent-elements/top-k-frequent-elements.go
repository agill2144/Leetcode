func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    mn := &minHeap{items: [][]int{}}
    for key, val := range freq {
        heap.Push(mn, []int{key, val})
        if mn.Len() > k {heap.Pop(mn)}
    }
    out := []int{}
    for i := 0; i < len(mn.items); i++ {
        out = append(out, mn.items[i][0])
    }
    return out
}


type minHeap struct {
	items [][]int
}

func (m *minHeap) Less(i, j int) bool {
	return m.items[i][1] < m.items[j][1]
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


// sort by freq
// tc = o(nlogn) + o(n)
// sc = o(n)
// func topKFrequent(nums []int, k int) []int {
//     freq := map[int]int{}
//     for i := 0; i < len(nums); i++ {freq[nums[i]]++}
//     sort.Slice(nums, func(i, j int)bool{
//         if freq[nums[i]] == freq[nums[j]] {
//             return nums[i] < nums[j]
//         }
//         return freq[nums[i]] < freq[nums[j]]
//     })
//     out := []int{}
//     for i := len(nums)-1; i >= 0 && len(out) != k; i-- {
//         lastNum := math.MinInt64
//         if len(out) != 0 {lastNum = out[len(out)-1]}
//         if nums[i] != lastNum {
//             out = append(out, nums[i])
//         }
//     }
//     return out
// }