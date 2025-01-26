func minCost(colors string, neededTime []int) int {
    total := 0
    prev := 0
    for i := 1; i < len(colors); i++ {
        if colors[i] == colors[prev] {
            if neededTime[i] < neededTime[prev] {
                // removing current, keeping prev
                total += neededTime[i]
            } else {
                // removing prev, keeping current
                total += neededTime[prev]
                prev = i
            } 
        } else {
            prev = i
        }
    }
    return total
}