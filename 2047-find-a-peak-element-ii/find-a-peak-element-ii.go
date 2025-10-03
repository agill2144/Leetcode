func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := m-1
    for left <= right {
        mid := left + (right-left)/2
        colIdx := 0
        midVal := mat[mid][0]
        for j := 1; j < n; j++ {
            if mat[mid][j] > midVal {
                midVal = mat[mid][j]
                colIdx = j
            }
        }
        above := math.MinInt64
        if mid-1 >= 0 {above = mat[mid-1][colIdx]}
        below := math.MinInt64
        if mid+1 < m {below = mat[mid+1][colIdx]}
        if midVal > above && midVal > below {
            return []int{mid, colIdx}
        }
        if above > midVal {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nil
}