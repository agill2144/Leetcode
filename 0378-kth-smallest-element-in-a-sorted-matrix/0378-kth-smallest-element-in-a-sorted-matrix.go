/*
    binary search on answers! the answer will lie in-betwen 0,0 and m-1,n-1
    binary search on values ( smallest value at 0,0 and largest value at m-1,n-1 )
    each mid is being evaluated, are you my kth value?
    for mid to be the kth value, our matrix must contain atleast k or more elements less than or equal to mid value
    if count of lessThanOrEqual to mid >= k, this is a potential kth value, however we want the smallest one, therefore search left ( we may have overshot )
    if count of lessThanOrEqual to mid < k, search right
    for mid to be a kth value, there must be atleast k elements that are <= mid

    time: o(log$maxNumber) for binary search on values * o(m+n) for stair case search
    space: o(1)

*/
func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    kth := -1    
    for left <= right {
        mid := left+(right-left)/2
        // for mid to be our kth smallest from the starting range
        // there must be atleast k elements to the left of this number
        // there must be atleast k elements less than or equal to this number
        numElementsLessThanMid := countLessThanOrEqualTo(matrix, mid)
        
        // this means that the number of elements to the left of mid is less than k
        // that means for sure, mid is not the kth value
        // for example; if mid = 8
        // and count = 2
        // start = 1
        // that means, from 1 to 8, there are only 2 elements in between (inclusive)
        // -------------------------------------------
        // 1  x  x  8
        // which means 8 currently is the 4th smallest .. and we want 8th smallest, therefore it must be on the right of 8
        // therefore search right
        if numElementsLessThanMid < k {
            left = mid+1
        } else {
            // we have ATLEAST k numbers to the left of this mid number, save this as a potential answer, and search left
            // we want the smallest possible such mid ( as this mid number may not even be a valid number number in our matrix )
            kth = mid
            right = mid-1
        }
    }
    return kth
}

// use stair case search to find count <= target
func countLessThanOrEqualTo(matrix [][]int,target int) int {
    m := len(matrix)
    n := len(matrix[0])
    count := 0
    r := m-1
    c := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            // all elements above this cell are less than or equal to target
            // therefore add all elements in this col
            count += r+1
            // and eval the next col
            c++
        } else {
            r--
        }
    }
    return count
}

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
// func kthSmallest(matrix [][]int, k int) int {
//     m := len(matrix)
//     n := len(matrix[0])
//     mx := &maxHeap{items: []int{}}
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             heap.Push(mx, matrix[i][j])
//             if mx.Len() > k {
//                 heap.Pop(mx)
//             }
//         }
//     }
//     return mx.items[0]
// }

// type maxHeap struct {
// 	items []int
// }
// func (m *maxHeap) Less(i, j int) bool {return m.items[i] > m.items[j]}
// func (m *maxHeap) Swap(i, j int) {m.items[i],m.items[j] = m.items[j], m.items[i]}
// func (m *maxHeap) Len() int {return len(m.items)}
// func (m *maxHeap) Push(x interface{}) { m.items = append(m.items, x.(int))}
// func (m *maxHeap) Pop() interface{} {
//     out := m.items[len(m.items)-1]
//     m.items = m.items[:len(m.items)-1]
//     return out
// }

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