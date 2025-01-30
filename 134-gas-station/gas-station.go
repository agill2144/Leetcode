func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    count := 0
    curr := 0
    for i := 0; i < 2*n; i++ {
        curr += gas[i%n]
        curr -= cost[i%n]
        if curr < 0 {
            count = 0
            curr = 0
        } else {
            count++
            if count == n {return (i+1)%n}
        }
    }
    return -1
}