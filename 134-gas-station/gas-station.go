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