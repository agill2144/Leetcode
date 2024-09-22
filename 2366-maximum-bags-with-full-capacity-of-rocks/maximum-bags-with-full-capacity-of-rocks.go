func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    diffs := []int{}
    for i := 0; i < len(capacity); i++ {
        diffs = append(diffs, capacity[i]-rocks[i])
    }
    sort.Ints(diffs)
    count := 0
    for i := 0; i < len(diffs); i++ {
        if diffs[i] == 0 {count++; continue}
        if additionalRocks >= diffs[i] {count++; additionalRocks-= diffs[i]}
    }
    return count
}