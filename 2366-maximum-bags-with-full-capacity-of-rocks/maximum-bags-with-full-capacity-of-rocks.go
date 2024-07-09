func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    n := len(capacity)
    diffs := make([]int, n)
    for i := 0; i < n; i++ {
        diffs[i] = capacity[i] - rocks[i]
    }
    sort.Ints(diffs)
    fulls := 0
    for i := 0; i < len(diffs); i++ {
        if diffs[i] == 0 {
            fulls++
            continue
        }
        needed := diffs[i]
        if additionalRocks >= needed {
            additionalRocks -= needed
            fulls++
        }
        if additionalRocks == 0 {break}
    }
    return fulls
}