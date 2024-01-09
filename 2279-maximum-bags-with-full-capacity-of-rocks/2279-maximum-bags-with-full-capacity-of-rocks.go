// greedy
// time = o(n) + o(nlogn) + o(n) ; o(nlogn)
// space = o(n)
func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    
    // 1. figure out the availibilites left
    avail := []int{}
    for i := 0; i < len(rocks); i++ { 
        avail = append(avail, capacity[i]-rocks[i])
    }
    // 2. sort them and fill up the bag that has most space (greedy)
    sort.Ints(avail)
    count := 0
    for i := 0; i < len(avail); i++ {
        // 3. if this bag is already full, count it and move on
        if avail[i] == 0 {
            count++
            continue
        }
        // 4. rocks > avail or rocks < avail ; in both cases, we want the min val
        // this is the number of rocks we can add to ith bag
        toBeAdded := min(avail[i], additionalRocks)
        // 5. if adding these many more rocks makes the bag full
        if avail[i]-toBeAdded == 0 {
            // 6. count the bag as full
            // and remove the number of rocks we just added from numberOfTotalRocks
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