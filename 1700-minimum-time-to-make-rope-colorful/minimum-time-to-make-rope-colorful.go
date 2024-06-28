/*
    greedy greedy greedy
    maintain 2 min times with a counter keeping track of consec
    1. smallest min
    2. second smallest min

    if this color is same as prev color
    - increment consecutive counter
    - set/update min and second min times
    - once we have seen 2 consecutive counts ( 2 same colors )
        - add min to total
        - promote second min to mind
        - reset second to maxInt
        - reset counter = 1
    
    time = o(n)
    space = o(1)
*/
func minCost(colors string, neededTime []int) int {
    total := 0
    min := math.MaxInt64
    second := math.MaxInt64
    count := 0
    for i := 0; i < len(colors); i++ {
        if i == 0 {
            count = 1
            min = neededTime[i]
            continue
        }
        if colors[i] == colors[i-1] {
            count++
            // update min and second min timers
            if neededTime[i] < min {
                second = min
                min = neededTime[i]
            } else if neededTime[i] < second {
                second = neededTime[i]
            }
            
            // have we seen 2 consecutive counts ?
            if count == 2 {
                count = 1
                total += min
                min = second
                second = math.MaxInt64
            }
        } else {
            // not consecutive, starting a new potential consecutive train
            // reset everything
            count = 1
            min = neededTime[i]
            second = math.MaxInt64
        }
    }
    return total
}

