func matrixMedian(grid [][]int) int {
    left := 1
    right := 1000000
    n := len(grid) * len(grid[0])
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := countOnLeft(grid, mid)
        if count >= (n/2)+1 {
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
        left := 0
        right := n-1
        // find the right most idx on left side of target ( target may exists, find the right most position)
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
        count += idx+1
    }
    return count
}
