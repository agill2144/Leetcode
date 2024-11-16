func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    n := len(rocks)
    diffs := []int{}
    total := 0
    for i := 0; i < n; i++ {
        diffs = append(diffs, capacity[i]-rocks[i])
    }
    sort.Ints(diffs)
    for i := 0; i < n && additionalRocks > 0 ; i++ {
        if diffs[i] == 0 ||  additionalRocks >= diffs[i] {
            if diffs[i] != 0 {additionalRocks -= diffs[i]}
            total++
        }
    }
    return total
}