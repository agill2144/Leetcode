func numberOfAlternatingGroups(colors []int, k int) int {
    total := 0
    n := len(colors)
    count := 1
    for i := 1; i < n+k-1; i++ {
        currIdx := i%n
        prevIdx := (i-1) % n
        if colors[currIdx] == colors[prevIdx] {
            count = 1
        } else {
            count++
        }
        if count == k {
            total++
            count--
        }
    }
    return total
}