func minCost(colors string, times []int) int {
    prev := 0
    time := 0
    n := len(colors)
    for i := 1; i < n; i++ {
        if colors[i] == colors[prev] {
            if times[prev] < times[i] {
                time += times[prev]
                prev = i
            } else {
                time += times[i]
            }
        } else {
            prev = i
        }
    }
    return time
}