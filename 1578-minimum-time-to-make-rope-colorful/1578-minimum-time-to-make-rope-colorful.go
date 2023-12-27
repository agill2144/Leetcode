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

