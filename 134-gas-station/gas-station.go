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
    totalGas := 0
    totalCost := 0
    for i := 0; i < n; i++ {
        totalGas += gas[i]
        totalCost += cost[i]
    }
    if totalGas < totalCost {return -1}

    total := 0
    result := 0
    for i := 0; i < n; i++ {
        total += (gas[i]-cost[i])
        if total < 0 {
            total = 0
            result = i+1
        }
    }
    return result
}