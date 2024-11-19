func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := m-1
    for left <= right {
        mid := left+(right-left)/2
        j := -1
        for i := 0; i < n; i++ {
            if j == -1 || mat[mid][i] > mat[mid][j] {j=i}
        }

        topVal := math.MinInt64
        if mid-1 >= 0 {topVal = mat[mid-1][j]}
        bottomVal := math.MinInt64
        if mid+1 < m {bottomVal = mat[mid+1][j]}
        if mat[mid][j] > topVal && mat[mid][j] > bottomVal {return []int{mid, j}}

        if topVal > bottomVal {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nil
}