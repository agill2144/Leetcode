func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2

        // find the max possible element in this col ( colIdx = mid )
        rowIdx := -1
        for i := 0; i < m; i++ {
            if rowIdx == -1 || mat[i][mid] > mat[rowIdx][mid] {
                rowIdx = i
            }
        }

        // now we have row and col(mid) idx, check if this cell is a peak
        leftCell := -1
        rightCell := -1
        if mid-1 >= 0 {leftCell = mat[rowIdx][mid-1]}
        if mid+1 < n  {rightCell = mat[rowIdx][mid+1]}
        val := mat[rowIdx][mid]
        if val > leftCell && val > rightCell {
            return []int{rowIdx, mid}
        }

        if rightCell > val {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return nil
}
/*
    approach:  brute force
    - go over entire matrix
    - and evaluate each cell
    time = o(mn)
    space = o(1)
*/
// func findPeakGrid(mat [][]int) []int {
//     m := len(mat)
//     n := len(mat[0])
//     dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             peak := true
//             for _, dir := range dirs {
//                 nr, nc := i+dir[0], j+dir[1]
//                 if nr >= 0 && nr < m && nc >= 0 && nc < n && mat[nr][nc] > mat[i][j] {
//                     peak = false
//                     break
//                 }
//             }
//             if peak {
//                 return []int{i, j}
//             }
//         }
//     }
//     return nil
// }