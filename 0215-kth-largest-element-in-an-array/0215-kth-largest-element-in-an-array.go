/*
    approach: quick select
    - like quick sort except you decide after each iteration, which side to sort on
*/

func findKthLargest(nums []int, k int) int {
    if k > len(nums) {return -1}
    targetIdx := len(nums)-k
    
    l := 0
    r := len(nums)-1
    for l <= r {
        
        nextSmaller := l
        pivotIdx := r
        for i := l; i < pivotIdx; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[i], nums[nextSmaller] = nums[nextSmaller], nums[i]
                nextSmaller++
            }
        }
        nums[nextSmaller], nums[pivotIdx] = nums[pivotIdx], nums[nextSmaller]
        if nextSmaller == targetIdx {
            return nums[nextSmaller]
        }
        
        if targetIdx > nextSmaller {
            l = nextSmaller+1
        } else {
            r = nextSmaller-1
        }
        
    }
    return -1
}

/*
    approach: min heap
    time: o(nlogk)
    space: o(k)
*/
// func findKthLargest(nums []int, k int) int {
//     mn := &minHeap{items: []int{}}
//     for i := 0; i < len(nums); i++ {
//         heap.Push(mn, nums[i])
//         if len(mn.items) > k {
//             heap.Pop(mn)
//         }
//     }
//     return mn.items[0]
// }

// type minHeap struct {
//     items []int
// }

// func (m *minHeap) Less(i, j int) bool {return m.items[i] < m.items[j]}
// func (m *minHeap) Swap(i, j int) {m.items[i],m.items[j] = m.items[j], m.items[i]}
// func (m *minHeap) Len() int {return len(m.items)}
// func (m *minHeap) Push(x interface{}) { m.items = append(m.items, x.(int))}
// func (m *minHeap) Pop() interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }

// approach: sort
// time: o(nlogn) 
    // - golang uses quickSort under the hood
// avg time: o(nlogn) and worst case is o(n^2) because quick sort
// space: o(n) used by quick sort
// func findKthLargest(nums []int, k int) int {
//     if k > len(nums) {return -1}
//     sort.Ints(nums)
//     return nums[len(nums)-k]
// }