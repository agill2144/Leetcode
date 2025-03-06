func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    freq := make([]int, (n*n)+1)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            freq[grid[i][j]]++
        }
    }
    out := make([]int, 2)
    for i := 1; i < len(freq); i++ {
        if freq[i] == 0 {
            out[1] = i
        } else if freq[i] > 1 {
            out[0] = i
        }
        if out[0] != 0 && out[1] != 0 {break}
    }
    return out
}