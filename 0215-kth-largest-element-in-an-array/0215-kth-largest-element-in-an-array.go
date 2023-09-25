/*
    approach: quick select
    - sort the side we care about
    - which side do we care about ? 
        - this depends on on the kth element
        - targetIdx = len(nums)-k
    - using a pivot idx, and a window ( left, right )
        - place pivot at the right ptr
        - and push all elements from left to pivotIdx (excluding pivotIdx)
        - to the left side
        -
*/
func findKthLargest(nums []int, k int) int {
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
    - use min heap to keep track of k elements
    time = o(nlogk)
    space = o(k)
    
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

/*
    brute force:
    - sort the array
    - return len(arr)-k idx element
    time = o(nlogn)
    space = maybe o(n) if sort is used mergeSort alg
    
    but what is k is larger than the input array
    - make k relative to the size of arry ; k = k % len(arr)
*/
// func findKthLargest(nums []int, k int) int {
//     if nums == nil || len(nums) == 0 {return -1}
//     if len(nums) == 1 {return nums[0]}
//     sort.Ints(nums)
//     return nums[len(nums)-k]
// }