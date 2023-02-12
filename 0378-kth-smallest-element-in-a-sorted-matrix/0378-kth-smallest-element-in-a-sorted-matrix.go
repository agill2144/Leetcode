func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    ans := -1    
    for left <= right {
        mid := left + (right-left)/2
        count := countLessOrEqual(mid, matrix)
        // there can be dupes, we want the smallest possible position that satifies count >= k
        // we may have overshot, so continue searching left.
        if count >= k {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}
func countLessOrEqual(target int, matrix[][]int) int {
    count := 0
    m := len(matrix)
    n := len(matrix[0])
    r := m-1
    c := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            count += r+1
            c++
        } else {
            r--
        }
    }
    return count
}