func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    curr := 0
    endIdx := -1
    for i := 0; i < 2*n; i++ {
        curr += gas[i%n]
        curr -= cost[i%n]
        if curr < 0 {endIdx = -1; curr = 0; continue}
        if i % n == endIdx {return endIdx}
        if endIdx == -1 {endIdx = i}
    }
    return -1
}