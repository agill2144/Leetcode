/*
    the goal is to take a full trip without running out of gas
    that is, we must visit all gas stations without running out of gas
    and we need to find which such idx position helps us achieve, such that
    if we start from this idx, we are able to take a full circular trip and end 
    up at this idx without running of gas

    - so lets evaluate each idx position
    - start with idx 0, fill in gas[0], burn cost[0]
        - if tank goes negative, this is not a good starting position
        - therefore start from next idx!
        - and so on...
    
    approach: 
*/
func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    goalIdx := 0
    tank := 0
    i := 0
    for i < 2*n {
        // fill up
        tank += gas[i%n]
        // reduce the travel cost
        tank -= cost[i%n]
        // can we proceed ?
        if tank >= 0 {
            i++
            if i%n == goalIdx {return goalIdx}
        } else {
            // no we cannot, stop, drop everything
            // and start new eval from next gas station
            i++
            tank = 0
            goalIdx = i%n
        }
    }
    return -1
}