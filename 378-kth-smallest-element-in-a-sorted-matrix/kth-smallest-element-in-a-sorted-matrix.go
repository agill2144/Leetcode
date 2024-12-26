// binary search on answers
func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    left := matrix[0][0]
    right := matrix[n-1][n-1]
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := countLessThanOrEqual(matrix, mid)
        if count >= k {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func countLessThanOrEqual(matrix [][]int, target int) int {
    n := len(matrix)
    count := 0
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