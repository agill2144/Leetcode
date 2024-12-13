// time = o(n) + o(nlogn) + o(n)
func findScore(nums []int) int64 {
    pairs := [][]int{} // [[idx,val]]
    for i := 0; i < len(nums); i++ {
        pairs = append(pairs, []int{i, nums[i]})
    }
    sort.Slice(pairs, func(i, j int)bool{
        if pairs[i][1] == pairs[j][1] {
            return pairs[i][0] < pairs[j][0]
        }
        return pairs[i][1] < pairs[j][1]
    })
    used := make([]bool, len(nums))
    var score int64 = 0
    for i := 0; i < len(pairs); i++ {
        idx, val := pairs[i][0], pairs[i][1]
        if used[idx] {continue}
        score += int64(val)
        if idx-1 >= 0 {used[idx-1] = true}
        if idx+1 < len(used){used[idx+1] = true}
    }
    return score
}


// min heap
// tc = o(nlogn) + o(nlogn)
// sc = o(n)
// func findScore(nums []int) int64 {
//     mn := &minHeap{items: [][]int{}}
//     for i := 0; i < len(nums); i++ {
//         heap.Push(mn, []int{i, nums[i]})
//     }
//     var score int64 = 0
//     used := make([]bool, len(nums))
//     for mn.Len() != 0 {
//         top := heap.Pop(mn).([]int)
//         idx, val := top[0], top[1]
//         if used[idx] {continue}
//         score += int64(val)
//         used[idx] = true
//         if idx-1 >= 0 {used[idx-1] = true}
//         if idx+1 < len(nums) {used[idx+1] = true}        
//     }
//     return score
// }

// type minHeap struct {
// 	items [][]int // [[idx, val]]
// }

// func (m *minHeap) Less(i, j int) bool {
// 	if m.items[i][1] == m.items[j][1] {
//         return m.items[i][0] < m.items[j][0]
//     }
//     return m.items[i][1] < m.items[j][1]
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
