func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    left := matrix[0][0]
    right := matrix[n-1][n-1]
    for left <= right {
        mid := left + (right-left)/2
        count := countOnLeft(matrix, mid)
        if count < k {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return right
    
}


func countOnLeft(matrix [][]int, target int) int {
    n := len(matrix)
    r := n-1
    c := 0
    total := 0
    for r >= 0 && c < n {
        if matrix[r][c] < target {
            total += r+1
            c++
        } else {
            r--
        }
    }
    return total
}
