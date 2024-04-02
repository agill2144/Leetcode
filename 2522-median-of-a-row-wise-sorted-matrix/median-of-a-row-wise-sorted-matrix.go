func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    total := m*n
    half := total/2

    left := 1
    right := 1000000
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := countOnLeft(grid, mid)
        if count > half {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}


func countOnLeft(grid [][]int, target int) int {
    m := len(grid)
    n := len(grid[0])
    count := 0
    for i := 0; i < m; i++ {
        idx := -1
        left := 0
        right := n-1
        for left <= right {
            mid := left + (right-left)/2
            midVal := grid[i][mid]
            if midVal <= target {
                idx = mid
                left = mid+1
            } else {
                right = mid-1
            }
        }
        count += (idx+1)
    }
    return count
}
