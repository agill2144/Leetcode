/*
    approach: min heap
    - use min heap to keep track of k elements
    time = o(nlogk)
    space = o(k)
    
*/
func findKthLargest(nums []int, k int) int {
    mn := &minHeap{items: []int{}}
    for i := 0; i < len(nums); i++ {
        heap.Push(mn, nums[i])
        if len(mn.items) > k {
            heap.Pop(mn)
        }
    }
    return mn.items[0]
}

type minHeap struct {
    items []int
}

func (m *minHeap) Less(i, j int) bool {return m.items[i] < m.items[j]}
func (m *minHeap) Swap(i, j int) {m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *minHeap) Len() int {return len(m.items)}
func (m *minHeap) Push(x interface{}) { m.items = append(m.items, x.(int))}
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

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