func minCost(costs [][]int) int {
    m := len(costs)
    n := len(costs[0])
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            costs[i][j] += minVal(i-1, j, costs)
        }
    }
    return min(costs[m-1][0], min(costs[m-1][1], costs[m-1][2]))
}

func minVal(row int, skipCol int, costs [][]int) int {
    out := math.MaxInt64
    for j := 0; j < len(costs[row]); j++ {
        if j == skipCol {continue}
        out = min(out, costs[row][j])
    }
    return out
}