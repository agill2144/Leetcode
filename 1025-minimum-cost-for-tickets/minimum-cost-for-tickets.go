// bottom up 1D dp
// time = o(lastDay); we run a linear loop from dayOne to lastDay
// space = o(lastDay); dp array
func mincostTickets(days []int, costs []int) int {
    lastDay := days[len(days)-1]
    dp := make([]int, lastDay+1)
    set := map[int]struct{}{}
    for i := 0; i < len(days); i++ {set[days[i]] = struct{}{}}
    dayOne := days[0]
    dp[dayOne] = min(costs[0], min(costs[1], costs[2]))
    for i := dayOne+1; i < len(dp); i++ {
        if _, ok := set[i]; !ok {dp[i] = dp[i-1]; continue}
        oneDayPassCost := costs[0] + dp[i-1]
        sevenDayPassCost := costs[1]
        if i-7 >= 0 {sevenDayPassCost += dp[i-7]}
        thirtyDayPassCost := costs[2]
        if i-30 >= 0 {thirtyDayPassCost += dp[i-30]}

        dp[i] = min(oneDayPassCost, min(sevenDayPassCost, thirtyDayPassCost))
    }
    return dp[len(dp)-1]
}

// brute force; dfs
// time = 3^n
// space = n
// func mincostTickets(days []int, costs []int) int {
//     ans := math.MaxInt64
//     oneDayPass := costs[0]
//     sevenDayPass := costs[1]
//     thirtyDayPass := costs[2]

//     var dfs func(start int, cost int)
//     dfs = func(start int, cost int) {
//         // base
//         if start >= len(days) {
//             ans = min(ans, cost)
//             return 
//         }

//         // logic
//         currDate := days[start]
//         dateAfterOneDay := -1
//         dateAfterSevenDays := -1
//         dateAfterThirtyDays := -1
//         for j := start+1; j < len(days); j++ {
//             newDate := days[j]
//             if (newDate >= currDate+1 && dateAfterOneDay == -1) {
//                 dateAfterOneDay = j
//             }
//             if newDate >= currDate+7 && dateAfterSevenDays == -1 {
//                 dateAfterSevenDays = j
//             }
//             if newDate >= currDate+30 && dateAfterThirtyDays == -1 {
//                 dateAfterThirtyDays = j
//             }
//         }
//         if dateAfterOneDay == -1 {dateAfterOneDay = len(days)}
//         if dateAfterSevenDays == -1 {dateAfterSevenDays = len(days)}
//         if dateAfterThirtyDays == -1 {dateAfterThirtyDays = len(days)}
//         dfs(dateAfterOneDay, cost+oneDayPass)
//         dfs(dateAfterSevenDays, cost+sevenDayPass)
//         dfs(dateAfterThirtyDays, cost+thirtyDayPass)
//     }
//     dfs(0,0)
//     return ans
// }
