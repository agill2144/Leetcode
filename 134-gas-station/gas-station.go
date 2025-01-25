func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    idx := -1
    curr := 0
    for i := 0; i < 2*n; i++ {
        if i%n == idx {return idx%n}
        if idx == -1 {idx = i%n}
        curr += gas[i%n]        
        curr -= cost[i%n]
        if curr < 0 {
            idx = -1
            curr = 0
        }
    }
    return -1
}