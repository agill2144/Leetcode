func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    total := m*n
    left := 0
    right := total-1
    for left <= right {
        mid := left + (right-left)/2
        r := mid / n
        c := mid % n
        if matrix[r][c] == target {return true}
        if target > matrix[r][c] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false
}