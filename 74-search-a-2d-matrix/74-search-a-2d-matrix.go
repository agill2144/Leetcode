func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    
    left := 0
    right := m*n-1
    
    for left <= right {
        mid := left + (right-left)/2
        r := mid/n
        c := mid%n
        if r >= m || c >= n {break}

        if matrix[r][c] == target {
            return true
        } else if matrix[r][c] > target {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return false
}