/*
    anything kth smallest/largest, consider heap
    or when sorting and want to trim dowm time on sort, consider heap
    
    approach: max heap
    - use a max heap to maintain the highest number at the top of the heap
    - when the heap size > k, drop the top
        - and we wouldnt care dropping the top because the top is going to be the highest number
        - and we want the smallest kth number
        - top of heap being highest possible number and bottom of heap being the smallest
        - protects our small elements, therefore max heap
    - loop over matrix and store each element in maxHeap
    - if maxHeap size grows bigger than k, drop the top of the heap
    - finally the top of the heap is the kth smallest
    
    time = o(mn * logK)
    space = o(K)
    
*/
func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    mx := &maxHeap{items: []int{}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            heap.Push(mx, matrix[i][j])
            if mx.Len() > k {
                heap.Pop(mx)
            }
        }
    }
    return mx.items[0]
}

type maxHeap struct {
	items []int
}
func (m *maxHeap) Less(i, j int) bool {return m.items[i] > m.items[j]}
func (m *maxHeap) Swap(i, j int) {m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Push(x interface{}) { m.items = append(m.items, x.(int))}
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

/*
    approach: brute force
    - created 1d array
    - sort 1d array
    - return k-1 element
    
    time = o(mn) + o(mnLogmn)
    space = o(mn)
*/
// func kthSmallest(matrix [][]int, k int) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     merged := []int{}
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             merged = append(merged, matrix[i][j])
//         }
//     }
//     sort.Ints(merged)
//     return merged[k-1]
// }