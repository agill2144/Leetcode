func minCostII(costs [][]int) int {
    m := len(costs)
    n := len(costs[0])
    
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            
            // we have n choices from above to pick for minimizing cost of this jth element
            cost := math.MaxInt64
            for k := 0; k < n; k++ {
                if k == j {continue}
                cost = min(cost, costs[i-1][k])
            }
            costs[i][j] += cost
            
        }
    }
    minCost := costs[m-1][0]
    for i := 1; i < n; i++ {
        minCost = min(minCost, costs[m-1][i])
    }
    return minCost
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}