func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    freq := make([]bool, (n*n)+1)
    out := make([]int, 2)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if freq[grid[i][j]] {
                out[0] = grid[i][j]
            } else {
                freq[grid[i][j]] = true
            }
        }
    }
    for i := 1; i < len(freq); i++ {
        if !freq[i] {out[1] = i; break}
    }
    return out
}
