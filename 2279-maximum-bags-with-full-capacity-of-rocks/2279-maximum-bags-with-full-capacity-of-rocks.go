func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    avail := []int{}
    for i := 0; i < len(rocks); i++ {
        avail = append(avail, capacity[i]-rocks[i])
    }
    sort.Ints(avail)
    count := 0
    for i := 0; i < len(avail); i++ {
        if avail[i] == 0 {
            count++
            continue
        }
        toBeAdded := min(avail[i], additionalRocks)
        if avail[i]-toBeAdded == 0 {
            count++
            additionalRocks -= toBeAdded
        }
    }
    return count
}

func min(x, y int) int {
    if x < y {return x}
    return y
}