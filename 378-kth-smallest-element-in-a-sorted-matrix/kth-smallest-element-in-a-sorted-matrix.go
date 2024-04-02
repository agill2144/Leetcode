func kthSmallest(matrix [][]int, k int) int {
    m := len(matrix)
    n := len(matrix[0])
    left := matrix[0][0]
    right := matrix[m-1][n-1]
    expectedOnLeft := k-1
    for left <= right {
        mid := left + (right-left)/2
        count := countOnLeft(matrix, mid)
        if count > expectedOnLeft {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return right
}

// strictly less than target ( not <= target )
func countOnLeft(matrix [][]int, target int ) int {
    m := len(matrix)
    n := len(matrix[0])
    r, c := m-1, 0
    count := 0
    for r >= 0 && c < n {
        if matrix[r][c] < target {
            count += (r+1)
            c++
        } else {
            r--
        }
    }
    // fmt.Println("target: ", target, "count: ", count)
    return count
}