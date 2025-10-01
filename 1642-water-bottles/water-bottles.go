func numWaterBottles(numBottles int, k int) int {
    full := numBottles
    empty := 0
    total := 0
    for full != 0 {
        // drink all
        total += full
        // get back
        getBack := (full + empty) / k
        // collect empty that could not be changed
        empty = (full + empty) % k
        full = getBack
    }
    return total
}