func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    left := matrix[0][0]
    right := matrix[n-1][n-1]
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        c := countLessOrEqual(matrix, mid)
        if c >= k {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func countLessOrEqual(matrix [][]int, target int) int {
    count := 0
    n := len(matrix)
    r := n-1
    c := 0
    for r >= 0 && c < n {
        if matrix[r][c] <= target {
            count += (r+1)
            c++
        } else {
            r--
        }
    }
    return count
}