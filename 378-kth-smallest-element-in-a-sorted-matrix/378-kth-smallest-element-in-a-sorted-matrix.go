
/*
    approach: brute force
    - Flatten the entire matrix in a 1D array
    - Sort it
    - Return the k-1 idx element
    Time: o(mn) + o(nlogn)
    Space: o(mn)
*/


// brute force
// func kthSmallest(matrix [][]int, k int) int {
//     flat := []int{}
//     m := len(matrix)
//     n := len(matrix[0])
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             flat = append(flat, matrix[i][j])
//         }
//     }
//     sort.Ints(flat)
//     return flat[k-1]
// }




/*
    approach: max heap of size k
    - For each element, store in maxHeap
    - If maxHeap size > k, pop to keep size <= k
    - Return the first element from maxHeap
    time : o(mn) x o(klogk)
    space: o(k)
*/


// max heap
func kthSmallest(matrix [][]int, k int) int {
    mx := &maxHeap{items: []int{}}
    m := len(matrix)
    n := len(matrix[0])
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
func (m *maxHeap) Swap(i, j int) { m.items[i],m.items[j] = m.items[j], m.items[i]}
func (m *maxHeap) Less(i, j int) bool {return m.items[i] > m.items[j]}
func (m *maxHeap) Len() int {return len(m.items)}
func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
func (m *maxHeap) Pop()interface{} {
	out := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return out
}