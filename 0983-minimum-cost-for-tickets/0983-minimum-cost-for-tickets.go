func mincostTickets(days []int, costs []int) int {
    dp := make([]int, days[len(days)-1]+1)
    
    oneDayPass := costs[0]
    sevenDayPass := costs[1]
    thirtyDayPass := costs[2]
    
    j := 0
    for i := 0; i < len(days); i++ { // o(n)
        // ensure j is at the same position as current day ( days[i] )
        for j != days[i] { // worst case input looks like [1,365] and we have ran a o(365 size loop) - which is constant!
            if j != 0 {
                dp[j] = dp[j-1]
            }
            j++
        }
        
        oneDayBack := 0
        sevenDaysBack := 0
        thirtyDaysBack := 0

        if j - 1 >= 0 { oneDayBack = dp[j-1]}
        if j - 7 >= 0 { sevenDaysBack = dp[j-7]}
        if j - 30 >= 0 { thirtyDaysBack = dp[j-30]}

        dp[j] = min(oneDayPass+oneDayBack, min(sevenDayPass+sevenDaysBack, thirtyDayPass+thirtyDaysBack))
        j++
           
    }
    // total time: o(n x 365) or o(n)
    // total space: o(365) or o(1)
    return dp[len(dp)-1]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}