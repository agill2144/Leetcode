func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        c := countOnLeft(matrix, mid)
        if c >= k {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func countOnLeft(matrix [][]int, target int) int {
    m := len(matrix)
    n := len(matrix[0])
    r := m-1; c := 0
    count := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            count += (r+1)
            c++
        } else {
            r--
        }
    }
    return count
}
// anything kth element = heap
// tc = o(mn * logk)
// sc = o(k)
// func kthSmallest(matrix [][]int, k int) int {
//     mx := &maxHeap{items: []int{}}
//     for i := 0; i < len(matrix); i++ {
//         for j := 0; j < len(matrix[i]); j++ {
//             heap.Push(mx, matrix[i][j])
//             if mx.Len() > k {heap.Pop(mx)}
//         }
//     }
//     return mx.items[0]
// }

// /*
// type Interface interface {
// 	sort.Interface
// 	Push(x any) // add x as element Len()
// 	Pop() any   // remove and return element Len() - 1.
// }
// */
// type maxHeap struct {
//     items []int
// }
// func (m *maxHeap) Len() int {return len(m.items)}
// func (m *maxHeap) Swap(i, j int) {m.items[i], m.items[j] = m.items[j], m.items[i]}
// func (m *maxHeap) Less(i, j int) bool {return m.items[i] > m.items[j]}
// func (m *maxHeap) Push(x interface{}) {m.items = append(m.items, x.(int))}
// func (m *maxHeap) Pop() interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }
