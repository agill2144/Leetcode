// sorting ? heaping ? bucket sorting ?
// quick select may also work
func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    deduped := []int{}
    for key, _ := range freq {deduped= append(deduped, key)}
    targetIdx := k-1
    left := 0
    right := len(deduped)-1
    for left <= right {
        nsf := left
        pivot := right
        for i := left; i < pivot; i++ {
            if freq[deduped[i]] >= freq[deduped[pivot]] {
                deduped[i], deduped[nsf] = deduped[nsf], deduped[i]
                nsf++
            }
        }
        deduped[pivot], deduped[nsf] = deduped[nsf], deduped[pivot]
        if nsf == targetIdx {return deduped[:k]}
        if targetIdx > nsf {
            left = nsf+1
        } else {
            right = nsf-1
        }
    }
    return nil
}
// sorting ? heaping?
// try bucket sort
// we want to sort by freq
// we can use bucket idxs are freq val
// and map each element to idx based on its freq ( freq == idx in bucket )
// because multiple elements could exist with same freq, 
// each idx in bucket should be able to save multiple elements ( because idx in bucket == freq val )
// therefore our bucket wont be a simple []int, it will be a [][]int
// bucket size is also known
// because we are sorting by freq, the largest possible idx we need in our bucket will be 
// the largest freq val, and that value will never exceed the size of nums array
// func topKFrequent(nums []int, k int) []int {
//     freq := map[int]int{}
//     for i := 0; i < len(nums); i++ {freq[nums[i]]++}
//     bucket := make([][]int, len(nums)+1)
//     for key, val := range freq {
//         bucket[val] = append(bucket[val], key)
//     }
//     out := []int{}
//     for i := len(bucket)-1; i >= 0 && len(out) != k; i-- {
//         if bucket[i] != nil {
//             out = append(out, bucket[i]...)
//         }
//     }
//     return out
// }

// tc = o(n) + o(nlogk) + o(k)
// sc = o(n) + o(k)
// func topKFrequent(nums []int, k int) []int {
//     freq := map[int]int{}
//     for i := 0; i < len(nums); i++ {freq[nums[i]]++}
//     mn := &minHeap{items: [][]int{}}
//     for key, val := range freq {
//         heap.Push(mn, []int{key, val})
//         if mn.Len() > k {heap.Pop(mn)}
//     }
//     out := []int{}
//     for i := 0; i < len(mn.items); i++ {
//         out = append(out, mn.items[i][0])
//     }
//     return out
// }


// type minHeap struct {
// 	items [][]int
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