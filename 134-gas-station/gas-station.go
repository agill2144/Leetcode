func canCompleteCircuit(gas []int, cost []int) int {
    idx := 0
    tank := 0
    n := len(gas)
    count := 0
    for i := 0; i < 2*n; i++ {
        // fill up
        tank += gas[i%n]
        // reduce the travel cost
        tank -= cost[i%n]
        // can we proceed ?
        if tank >= 0 {
            // yes, we have fuel to go to next
            count++
            if count == n {return idx}
        } else {
            // no we cannot, stop, drop everything
            // and start new eval from next gas station
            tank = 0
            count = 0
            idx = (i+1)%n
        }
    }
    return -1
}