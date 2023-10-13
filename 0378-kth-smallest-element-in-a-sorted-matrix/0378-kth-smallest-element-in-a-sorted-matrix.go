func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if countLessThanOrEqualTo(matrix, mid) < k {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}

// use stair case search to find count <= target
func countLessThanOrEqualTo(matrix [][]int,target int) int {
    m := len(matrix)
    n := len(matrix[0])
    count := 0
    r := m-1
    c := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            // all elements above this cell are less than or equal to target
            // therefore add all elements in this col
            count += r+1
            // and eval the next col
            c++
        } else {
            r--
        }
    }
    return count
}