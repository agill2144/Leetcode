func findPeakGrid(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    left := 0
    right := m-1
    // pick a row using bs
    for left <= right {
        mid := left + (right-left)/2
        
        // find the max in this row
        maxIdx := 0
        for j := 0; j < n; j++ {
            if matrix[mid][j] > matrix[mid][maxIdx] {maxIdx = j}
        }

        // compare with above and below values
        top := math.MinInt64
        if mid-1 >= 0 {top = matrix[mid-1][maxIdx]}
        bottom := math.MinInt64
        if mid+1 < m {bottom = matrix[mid+1][maxIdx]}
        curr := matrix[mid][maxIdx]
        if curr > top && curr > bottom {return []int{mid, maxIdx}}

        if top > curr {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nil
}