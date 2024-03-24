// bottom up 1D dp
// time = o(lastDay); we run a linear loop from dayOne to lastDay
// space = o(lastDay); dp array
func mincostTickets(days []int, costs []int) int {
    lastDay := days[len(days)-1]
    dp := make([]int, lastDay+1)

    daysIdx := 0
    for i := 1; i < len(dp); i++ {
        // not a day we are going out on
        if i < days[daysIdx] {
            dp[i] = dp[i-1]
            continue
        }
        // is a day we are going out on
        one := costs[0]
        seven := costs[1]
        thirty := costs[2]
        if i-1 >= 0 {one += dp[i-1]}
        if i-7 >= 0 {seven += dp[i-7]}
        if i-30 >= 0 {thirty += dp[i-30]}
        dp[i] = min(one, min(seven, thirty))
        daysIdx++
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