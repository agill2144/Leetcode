func findPeakGrid(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])

    left := 0
    right := m-1

    // we binary search on m rows
    // time = o(logm)
    for left <= right {
        mid := left + (right-left)/2
        maxCol := math.MinInt64
        maxColIdx := -1

        // for each row, we loop over n cols
        // time = o(n)
        for j := 0; j < n; j++ {
            if matrix[mid][j] > maxCol {
                maxCol = matrix[mid][j]
                maxColIdx = j
            }
        }

        // check if matrix[mid][maxColIdx] is peak
        // we only need to check top and bottom cells
        // because we picked max elemenet in this row
        // therefore left and right neigbors check is already done
        topVal := math.MinInt64
        if mid-1 >= 0 {topVal = matrix[mid-1][maxColIdx]}
        bottomVal := math.MinInt64
        if mid+1 < m {bottomVal = matrix[mid+1][maxColIdx]}

        if maxCol > topVal && maxCol > bottomVal {
            return []int{mid, maxColIdx}
        }

        if topVal > maxCol {
            // take binary search on row above curr ($mid) row
            right = mid-1
        } else {
            left = mid+1
        }
    }
    // total time = o(logm * n)
    // space = o(1)
    return nil
}