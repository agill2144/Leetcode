/*
approach: bucket sort / counting sort
- use map as a bucket instead of an array
- why ? we have negative values, and we cannot easily plot/map a negative value to a idx in an array
- therefore using map
- time = o(n) + o(max-min+1) 
- space = o(n)
*/
func findKthLargest(nums []int, k int) int {
    n := len(nums)
    targetIdx := n-k
    freq := map[int]int{}
    start, end := math.MaxInt64, math.MinInt64
    for i := 0; i < n; i++ {
        start = min(start, nums[i])
        end = max(end, nums[i])
        freq[nums[i]]++
    }

    idx := -1
    for i := start; i <= end; i++ {
        val, ok := freq[i]
        if !ok {continue}
        idx += val
        if idx >= targetIdx {return i}
    }
    return -1
}
// func findKthLargest(nums []int, k int) int {
//     n := len(nums)
//     targetIdx := n-k
//     left := 0
//     right := n-1
//     for left <= right {
//         pivot := right
//         ns := left
//         for i := left; i < pivot; i++ {
//             if nums[i] < nums[pivot] {
//                 nums[i], nums[ns] = nums[ns], nums[i]
//                 ns++
//             }
//         }
//         nums[ns], nums[pivot] = nums[pivot], nums[ns]
//         if ns == targetIdx {return nums[ns]}
//         if targetIdx < ns {
//             right = ns - 1
//         } else {
//             left = ns + 1
//         }
//     }
//     return -1
// }
// func findKthLargest(nums []int, k int) int {
//     n := len(nums)
//     mn := &minHeap{items: []int{}}
//     // o(n)
//     for i := 0; i < n; i++ {
//         // o(logk)
//         heap.Push(mn, nums[i])
//         if mn.Len() > k {
//             // o(logk)
//             heap.Pop(mn)
//         }
//     }
//     // time = o(n * 2logk) = o(nlogk)
//     // space = o(k) 
//     return mn.items[0]
// }


// type minHeap struct {
// 	items []int
// }

// func (m *minHeap) Less(i, j int) bool {
// 	return m.items[i] < m.items[j]
// }
// func (m *minHeap) Swap(i, j int) {
// 	m.items[i], m.items[j] = m.items[j], m.items[i]
// }
// func (m *minHeap) Len() int {
// 	return len(m.items)
// }
// func (m *minHeap) Push(x interface{}) {
// 	m.items = append(m.items, x.(int))
// }
// func (m *minHeap) Pop() interface{} {
// 	out := m.items[len(m.items)-1]
// 	m.items = m.items[:len(m.items)-1]
// 	return out
// }
