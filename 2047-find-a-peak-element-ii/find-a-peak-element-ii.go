func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := m-1
    for left <= right {
        mid := left + (right-left)/2
        idx := 0
        for j := 1; j < n; j++ {
            if mat[mid][j] > mat[mid][idx] {
                idx = j
            }
        }
        curr := mat[mid][idx]
        top := math.MinInt64
        if mid-1 >= 0 {top = mat[mid-1][idx]}
        bottom := math.MinInt64
        if mid+1 < m {bottom = mat[mid+1][idx]}
        if curr > top && curr > bottom {return []int{mid, idx}}
        if top > bottom {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return nil
}