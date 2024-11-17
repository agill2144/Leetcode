func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := m-1
    for left <= right {
        mid := left + (right - left)/2
        
        maxColIdx := -1
        for j := 0; j < n; j++ {
            if maxColIdx == -1 || mat[mid][j] > mat[mid][maxColIdx] {
                maxColIdx = j
            }
        }

        topVal := math.MinInt64
        if mid-1 >= 0 {topVal = mat[mid-1][maxColIdx]}
        bottomVal := math.MinInt64
        if mid+1 < m {bottomVal = mat[mid+1][maxColIdx]}
        if mat[mid][maxColIdx] > topVal && mat[mid][maxColIdx] > bottomVal {
            return []int{mid, maxColIdx}
        }
        if topVal > bottomVal {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nil
}