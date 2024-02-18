/*
*/
func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    total := m*n
    half := total/2
    left := 0
    right := 1000002
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < m; i++ {
            count += countOnLeft(grid[i], mid)
        }
        if count <= half {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return left
}

func countOnLeft(arr []int, target int) int {
    left := 0
    right := len(arr)-1
    // find the right most position of target on left of target
    idx := -1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] <= target {
            idx = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx+1
}
/*
    approach: brute force
    - create a 1d array
    - sort it
    - return the middle
    time = o(mn) + o(mn log mn)
    space = o(mn)
*/
// func matrixMedian(grid [][]int) int {
//     merged := []int{}
//     m := len(grid)
//     n := len(grid[0])
//     mid := (m*n)/2
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             merged = append(merged, grid[i][j])
//         }
//     }
//     sort.Ints(merged)
//     return merged[mid]
// }