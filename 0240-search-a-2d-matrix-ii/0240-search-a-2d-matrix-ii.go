/*
    approach: brute force
    - flatten the matrix into a 1D array
    - Sort the 1D array
    - Binary Search on 1D array
    time: 
        o(mn) + o(mnlogmn) + o(logmn)
    space: 
        o(mn)
    
    approach: suboptimal
    - for each row
    - do a binary search
    time: 
        o(mlogn)
    space: 
        o(1)
    
    approach: stair case search
    - anytime you see coloums are sorted and rows are sorted, consider staircase search
    - We can start from any of the 2 corners
        - Bottom left
        - Top Right
        - Why only these 2 ?
        - Because if you look carefully, these are the only 2 corners from which we can definatively pick a direction
            - ex: bottom-left : we can either go right if we are looking for a larger number, or go up if we are looking for a smaller number
            - ex: top-right: we can either go left if we are looking for a smaller number , or go down if we are looking for a larger number
            - We cannot definatively make a choice from top left or bottom right corner
    - The idea is to pick a corner out of the above 2 corners ^
    - Then checking if this cell is the target
    - if it is, return true
    - otherwise pick a direction based on how target compares with current cell value and based on the corner choosen to start from.
    
    time: o(m+n)
    space: o(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    r := m-1
    c := 0
    
    for r >= 0 && r < m && c >= 0 && c < n {
        if matrix[r][c] == target {
            return true
        } else if target > matrix[r][c] {
            c++
        } else {
            r--
        }
    }
    return false
}