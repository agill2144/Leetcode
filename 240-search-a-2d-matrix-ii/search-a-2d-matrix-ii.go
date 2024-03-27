/*
    "row are sorted in ascending from left to right"
    "column are sorted in ascending from top to bottom"

    approach: staircase search
    - pick a corner ( bottom left or top right )
    - and check if its our target
    - if not , we know a definitive direction to proceed into to look for our target
    - keep doing this while we are still inbound

    time = o(m+n)
    space = o(1)   
*/
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    r, c := m-1, 0

    for r >= 0 && c < n {
        if matrix[r][c] == target {return true}
        if target > matrix[r][c] {
            c++
        } else {
            r--
        }
    }
    return false
}