func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    left := matrix[0][0]
    right := matrix[n-1][n-1]
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        count := countLessThanOrEqualTo(matrix, mid)
        if count >= k {
            res = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return res
}

func countLessThanOrEqualTo(matrix [][]int, target int) int {
    n := len(matrix)
    r := n-1
    c := 0
    count := 0
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