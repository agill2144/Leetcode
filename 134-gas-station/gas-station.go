func canCompleteCircuit(gas []int, cost []int) int {
    curr := 0
    n := len(gas)
    count := 0
    start := 0
    for i := 0; i < 2*n; i++ {
        
        // fill up
        curr += gas[i%n]
        // travel to next
        curr -= cost[i%n]
        
        // if we are out (< 0), we reset
        if curr < 0 {
            curr = 0
            count = 0
            start = (i+1)%n
        } else {
            count++
            if count == n {return start}
        }
    }
    if count != n {return -1}
    return start    
}