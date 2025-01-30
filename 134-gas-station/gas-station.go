func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    destIdx := -1
    curr := 0
    for i := 0; i < 2*n; i++ {
        if (i%n) == destIdx {return destIdx}
        if destIdx == -1 {destIdx = i%n}
        curr += gas[i%n]
        curr -= cost[i%n]
        if curr < 0 {
            destIdx = -1
            curr = 0
        }
    }
    return -1
}