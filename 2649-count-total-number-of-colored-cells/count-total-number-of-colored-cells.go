func coloredCells(n int) int64 {
    x := 1
    total := 1
    for x != n {
        total += (4*x)
        x++
    }
    return int64(total)
}