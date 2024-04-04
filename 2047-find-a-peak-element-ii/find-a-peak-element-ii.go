/*
    - similar to finding peak element in a 1D array
    - a cell is a peak element, if its bigger than its 4 adj neighbors
    - in the 1d array, we used binary search to find the peak element
    - how can we use binary search in this matrix to find peak element?
    - we WILL keep the same intuition as finding peak in a 1D array
    - if mid element is not peak, take binary search towards the higher element
    
    approach: binary search on cols
    - binary search on cols, randomly
    - use binary search mid idx, to evaluate a col
    - we want peak element from this col, and for each cell
        - we have to check 4 adj cells to check if its peak
    - instead of checking each cell, 
    - our best bet is to only evaluate the max number present in this col
    - which means, find the max element in this col
    - finding and use max element from a col, takes care of 2 adj neighbors automatically
        - top and bottom
    - NOW this becomes a plain "finding peak in a 1D array"
    - now we have a cell to check
        - compare this cell with its left and right adj cells
        - if its peak, return the row=maxElementIdx, and col=mid
    - if a cell is not peak, pick a side / cell which has higher value COMPARED TO MID
        - remember mid = col idx
    - if right cell has higher value compared to mid cell, take search on the right side
    - otherwise take the search on the left side

    - we are doing a binary search on n cols
        - o(logn)
    - and at each cell, we look for a max element in m rows
        - o(m)
    time = o(logn * m)
    space = o(1)
*/
func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        row := 0
        for i := 1; i < m; i++ {
            if mat[i][mid] > mat[row][mid] {
                row = i
            }
        }
        // max element in mid col is at row,mid(col)
        maxElement := mat[row][mid]
        leftNei, rightNei := -1,-1
        if mid-1 >= 0 {leftNei = mat[row][mid-1]}
        if mid+1 < n {rightNei = mat[row][mid+1]}
        if maxElement > leftNei && maxElement > rightNei {
            return []int{row, mid}
        }
        if leftNei > maxElement {
            right = mid-1
        } else {
            left = mid+1
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