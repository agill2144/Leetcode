func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    count := 0
    curr := 0
    for i := 0; i < 2*n; i++ {
        curr += gas[i%n]
        curr -= cost[i%n]
        if curr >= 0 {
            count++
            if count == n {return (i+1)%n}
        } else {
            count = 0
            curr = 0
        }
    }
    return -1
}
// func canCompleteCircuit(gas []int, cost []int) int {
//     n := len(gas)
//     curr := 0
//     endIdx := -1
//     for i := 0; i < 2*n; i++ {
//         if i % n == endIdx {return endIdx}
//         curr += gas[i%n]
//         curr -= cost[i%n]
//         if endIdx == -1 {endIdx = i}
//         if curr < 0 {endIdx = -1; curr = 0; continue}
       
//     }
//     return -1
// }