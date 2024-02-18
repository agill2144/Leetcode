/*
    approach: brute force
    - create a 1d array
    - sort it
    - return the middle
*/
func matrixMedian(grid [][]int) int {
    merged := []int{}
    m := len(grid)
    n := len(grid[0])
    mid := (m*n)/2
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            merged = append(merged, grid[i][j])
        }
    }
    sort.Ints(merged)
    return merged[mid]
}