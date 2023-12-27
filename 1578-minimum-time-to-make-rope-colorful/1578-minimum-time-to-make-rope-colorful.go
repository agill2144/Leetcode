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
    second := math.MaxInt64
    min := neededTime[0]
    count := 1
    for i := 1; i < len(colors); i++ {
        if colors[i] == colors[i-1] {
            count++
            if neededTime[i] < min {
                second = min
                min = neededTime[i]
            } else if neededTime[i] < second {
                second = neededTime[i]
            }
            if count == 2 {
                count = 1
                total += min
                min = second
                second = math.MaxInt64
            }
        } else {
            count = 1
            min = neededTime[i]
            second = math.MaxInt64
        }
    }
    return total
}

