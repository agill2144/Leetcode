func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    total := m*n
    expectedCount := (total/2)+1
    left := 1
    right := 1000000
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        if count(grid, mid) >= expectedCount {
            res = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return res
}

// count num of elements <= target
func count(grid [][]int, target int) int {
    c := 0
    m := len(grid)
    n := len(grid[0])
    for i := 0; i < m; i++ {
        left := 0
        right := n-1
        idx := -1
        for left <= right {
            mid := left + (right-left)/2
            if grid[i][mid] <= target {
                idx = mid
                left = mid+1
            } else {
                right = mid-1
            }
        }
        c += (idx+1)
    }
    return c
}