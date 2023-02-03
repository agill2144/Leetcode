/*
    approach: quick select
    - pick a pivot idx to compare all elements with
    - lets make pivot as last idx
    - we need to ensure all elements less than pivot value is to the left
    - so we need a nextSmaller ptr that collects all the smaller value than pivot
    - the nextSmaller ptr becomes responsible for collecting all smaller values than pivot
    - each time we run into a value thats smaller than pivot value, we swap with nextSmaller
        - move nextSmaller forward
    - once we have done that loop once, we know for sure all elements to the left is smaller
    - nextSmaller may not reach pivotIdx-1, it may stop early
    - this tells us that the pivot value belongs ( 100% ) at the nextSmaller idx
        - therefore swap the nextSmaller with pivot idx once the we have done a loop
    - now we know for sure that the value at nextSmaller is at the correct idx position
        - in terms of asc order
    - then we can check if nextSmaller is the targetIdx ( n - k )
        - if yes, return and exit earl
    - if targetIdx lies on the right side,
        - run the same quickSelect from nextSmaller+1 : end
        - we exclude nextSmaller because we have already checked its not our target idx
        - we exclude nextSmaller because we have it at the correct position, therefore no need to include it again

    - otherwise
        - run the same quickSelect from start : nextSmaller-1
        - we exclude nextSmaller because we have already checked its not our target idx
        - we exclude nextSmaller because we have it at the correct position, therefore no need to include it again
        
    avg time: o(n)
    worst time: o(n^2)
    space: o(1) in iterative
*/

// func findKthLargest(nums []int, k int) int {
//     if k > len(nums) {return -1}
//     targetIdx := len(nums)-k
    
//     l := 0
//     r := len(nums)-1
//     for l <= r {
        
//         nextSmaller := l
//         pivotIdx := r
//         for i := l; i < pivotIdx; i++ {
//             if nums[i] < nums[pivotIdx] {
//                 nums[i], nums[nextSmaller] = nums[nextSmaller], nums[i]
//                 nextSmaller++
//             }
//         }
//         nums[nextSmaller], nums[pivotIdx] = nums[pivotIdx], nums[nextSmaller]
//         if nextSmaller == targetIdx {
//             return nums[nextSmaller]
//         }
        
//         if targetIdx > nextSmaller {
//             l = nextSmaller+1
//         } else {
//             r = nextSmaller-1
//         }
        
//     }
//     return -1
// }

// quick select recursive
func findKthLargest(nums []int, k int) int {
    if k > len(nums) {return -1}
    l := 0
    r := len(nums)-1
    targetIdx := len(nums)-k
    
    var quickSelect func(left, right int) (int, bool)
    quickSelect = func(left, right int) (int, bool) {
        // base
        if l > r {return -1, false}
        
        // logic
        // pivot is the last idx (i.e right idx)
        pivotIdx := right
        nextSmaller := left
        
        
        // from left - exlcuding pivotIdx
        for i := left; i < pivotIdx; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[i], nums[nextSmaller] = nums[nextSmaller], nums[i]
                nextSmaller++
            }
        }
        // we found the 100% correct position for pivot val ( i.e nextSmaller )
        nums[nextSmaller], nums[pivotIdx] = nums[pivotIdx], nums[nextSmaller]

        // if nextSmaller is our kth largest
        if nextSmaller == targetIdx {
            // return
            return nums[nextSmaller], true
        } else if targetIdx > nextSmaller {
            if val, ok := quickSelect(nextSmaller+1, right); ok {return val, ok}
        }
        return quickSelect(left, nextSmaller-1)
    }
    val, _ := quickSelect(0, len(nums)-1)
    return val
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