func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    total := m*n
    median := total/2

    left := 0
    right := 1000002
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < m; i++ {
            count += countOnLeft(grid[i], mid)
        }
        if count > median {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return right
}


func countOnLeft(arr []int, target int) int {
    left := 0
    right := len(arr)-1
    // find the right most position of target on left of target
    idx := -1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] < target {
            idx = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx+1
}
