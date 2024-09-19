func canCompleteCircuit(gas []int, cost []int) int {
    start := 0
    tank := 0
    count := 0
    n := len(gas)
    for i := 0; i < 2 * n; i++ {
        tank += gas[i%n]
        tank -= cost[i%n]
        if tank >= 0 {
            count++
            if count == n {return start}
        } else {
            tank = 0
            count = 0
            start = (i+1) % n
        }
    }
    return -1
}