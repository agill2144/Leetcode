func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    
    for left < right {
        mid := left + (right-left)/2
        count := countLessOrEqual(mid, matrix)
        if count < k {
            left = mid+1
        } else {
            right = mid
        }
    }
    return left
}
func countLessOrEqual(target int, matrix[][]int) int {
    count := 0
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            if matrix[i][j] <= target {count++}
        }
    }
    return count
}