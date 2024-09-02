func chalkReplacer(chalk []int, k int) int {
    total := 0
    for i := 0; i < len(chalk); i++ {total += chalk[i]; if total > k {break}}
    k %= total
    idx := 0
    n := len(chalk)
    for k >= chalk[idx%n] {
        if k >= chalk[idx%n] {
            k -= chalk[idx%n]
            idx++
        }
    }
    return idx%n
}