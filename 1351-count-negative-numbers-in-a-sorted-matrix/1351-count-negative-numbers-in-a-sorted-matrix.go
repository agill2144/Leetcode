/*
    This is a classic stair case search
    - Outline the staircase that holds all negatives
    
    time: o(m+n)
    space: o(1)
    
    
*/
func countNegatives(grid [][]int) int {   
    
    m := len(grid)
    n := len(grid[0])
    r := m-1
    c := 0
    totalNegatives := 0
    
    for r >= 0 && c < n {
        if grid[r][c] < 0 {
            totalNegatives += n-c
            r--
        } else {
            c++
        }
    }
    return totalNegatives
    
}