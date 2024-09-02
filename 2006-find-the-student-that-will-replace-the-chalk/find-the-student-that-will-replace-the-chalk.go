func chalkReplacer(chalk []int, k int) int {
    total := 0
    for i := 0; i < len(chalk); i++ {total += chalk[i]}
    k %= total
    idx := 0
    n := len(chalk)
    for true {
        if k >= total {
            k -= total
        } else if k >= chalk[idx%n] {
            k -= chalk[idx%n]
            idx++
        } else {
            break
        }
    }
    return idx%n
}