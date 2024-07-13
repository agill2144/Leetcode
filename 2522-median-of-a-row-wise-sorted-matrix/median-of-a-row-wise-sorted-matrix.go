func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    left := 1
    right := 1000000
    total := m*n
    half := total/2
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
    count := 0
    m := len(grid)
    n := len(grid[0])
    for i := 0; i < m; i++ {
        // because each row is sorted, we need to find the closest/last idx to target
        // from the left side of target
        //                           closestIdx
        // |----------------------------------|--|----------------------|
        //                                       Target
        // how will idx tell us the count ?
        // count = idx+1
        idx := -1
        left := 0
        right := n-1
        for left <= right {
            mid := left + (right-left)/2
            if grid[i][mid] <= target {
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
