func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    rem := []int{}
    for i := 0; i < len(capacity); i++ {
        diff := capacity[i] - rocks[i]
        rem = append(rem, diff)
    }
    sort.Ints(rem)
    total := 0
    for i := 0; i < len(rem); i++ {
        if rem[i] == 0 {total++; continue}
        if additionalRocks >= rem[i] {
            additionalRocks -= rem[i]
            total++
        } else {
            break
        }
    }
    return total
}