func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := m-1
    for left <= right {
        mid := left + (right-left)/2
        // find max in this row
        maxVal := mat[mid][0]
        maxValIdx := 0
        for j := 1; j < n; j++ {
            if mat[mid][j] > maxVal {
                maxVal = mat[mid][j]
                maxValIdx = j
            }
        }

        // now compare with above and bottom cell
        above := math.MinInt64
        if mid-1 >= 0 {above = mat[mid-1][maxValIdx]}
        bottom := math.MinInt64
        if mid+1 < m {bottom = mat[mid+1][maxValIdx]}
        curr := maxVal
        if curr > above && curr > bottom {
            return []int{mid, maxValIdx}
        }
        if above > curr {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nil
}