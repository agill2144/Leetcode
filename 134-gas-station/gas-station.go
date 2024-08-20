// like kadanes max subarray sum algorithm
// "start a chain / subarray "
// "if a chain / subarray goes negative, start a fresh one"
// similarly, evaluate starting at EACH gas station
// starting from 0th
// fuel up, travel and remove travel dist (cost[i])
// if we run out of gas ( current gas goes negative )
// it means that the idx we started from is not good
// therefore start a new evaluation from current idx
// what happens if the ans is one of the gas station we went over when our current gas had not went negative yet
// this is why we will run a 2*n loop, because there is gauranteed to be 1 answer or no answer
// time = o(2n)
// space = o(1)
func canCompleteCircuit(gas []int, cost []int) int {
    curr := 0
    n := len(gas)
    count := 0
    start := 0
    for i := 0; i < 2*n; i++ {
        
        // fill up
        if curr == 0 && count == 0{start = i%n}
        curr += gas[i%n]
        // travel to next
        curr -= cost[i%n]
        
        // if we are out (< 0), we reset
        // just like max sum subarray: "it means our subarray sum went negative"
        // there is no point in continuing in that subarray
        // start a new one!
        if curr < 0 {
            curr = 0
            count = 0
        } else {
            count++
            // if we were able to reach all n stations, then we can call it success!
            // the start of the subarray/chain is the correct gas station
            if count == n {return start}
        }
    }
    if count != n {return -1}
    return start    
}