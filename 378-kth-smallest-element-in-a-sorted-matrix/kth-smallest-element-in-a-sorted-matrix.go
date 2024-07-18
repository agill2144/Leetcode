func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    ans := -1
    tmpK := k-1
    for left <= right {
        mid := left + (right-left)/2
        count := countLessThanOrEqualTo(matrix, mid)
        if count > tmpK {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

func countLessThanOrEqualTo(matrix [][]int, target int) int {
    m := len(matrix)
    n := len(matrix[0])
    count := 0
    r := m-1
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
