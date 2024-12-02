func matrixMedian(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    left := 1
    right := 1000000
    totalEle := m*n
    minCount := (totalEle/2)+1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := countLessThanOrEqualTo(matrix, mid)
        // count <= median value
        // MUST ALWAYS BE midIdx+1 ( at the very least! )
        // therefore if count < minCount
        // search right
        if count < minCount {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}

func countLessThanOrEqualTo(matrix [][]int, target int) int {
    total := 0
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        idx := -1
        left := 0
        right := n-1
        for left <= right {
            mid := left + (right-left)/2
            if matrix[i][mid] <= target {
                idx = mid
                left = mid+1
            } else {
                right = mid-1
            }
        }
        total += idx+1
    }
    return total
}