/*
    - in an uniq sorted odd sized array
    - median is at (n-1)/2 idx
    - count num of elements <= median val will always be medianIdx+1

    - but in a sorted odd size array with dupes
    - median idx is still at (n-1)/2 idx
    - we want the keep searching for a value whose count num of elems <= itself >= medianIdx+1
*/
func matrixMedian(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    left := 1
    right := 1000000
    totalEles := m*n
    medianIdx := (totalEles-1)/2
    expectedCount := medianIdx+1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := countLessThanOrEqualTo(grid, mid)
        if count >= expectedCount {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func countLessThanOrEqualTo(grid [][]int, target int) int {
    m := len(grid)
    n := len(grid[0])
    total := 0
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
        total += (idx+1)
    }
    return total
}