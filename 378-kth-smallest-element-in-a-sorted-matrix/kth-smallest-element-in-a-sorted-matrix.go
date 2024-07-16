func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    ans := -1
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    for left <= right {

        mid := left + (right-left)/2
        count := countLessThanOrEqualTo(matrix, mid)
        // is mid our kth smallest ?
        // only if count on left of mid >= k elements
        // why > ?
        // because if mid is our kth smallest
        // then we have included this kth element in our counting (because count uses <= operator )
        // therefore if mid is our kth element, its count will ALWAYS be >= k elements on left
        // therefore save this potential answer and continue searching on left
        if count >= k {
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