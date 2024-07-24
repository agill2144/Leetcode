func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2


        maxRowIdx := -1
        maxInThisCol := math.MinInt64
        for i := 0; i < m; i++ {
            if mat[i][mid] > maxInThisCol{
                maxInThisCol = mat[i][mid]
                maxRowIdx = i
            }
        }

        if (mid == 0 || mat[maxRowIdx][mid] > mat[maxRowIdx][mid-1]) && (mid == n-1 || mat[maxRowIdx][mid] > mat[maxRowIdx][mid+1]) {
            return []int{maxRowIdx, mid}
        } else if (mid == n-1 || mat[maxRowIdx][mid+1] > mat[maxRowIdx][mid]) {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return nil
}