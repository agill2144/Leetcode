func minDays(bloomDay []int, m int, k int) int {
    left := 1
    right := -1
    for i := 0; i < len(bloomDay); i++ { right = max(right, bloomDay[i]) }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        // how many bouquets can be made if today was $mid day
        flowers := 0
        bouquets := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                flowers++
                if flowers == k {
                    bouquets++
                    flowers = 0
                }
            } else {
                flowers = 0
            }
        }

        // when does mid not work?
        // when number of bouquets made were < required m
        if bouquets < m {
            // we are too early, wait for next set of days
            // take search to right
            left = mid+1            
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans

}
// func minDays(bloomDay []int, m int, k int) int {
//     /*
//         - we need to make m bouquets
//         - for 1 bouquets, we need k flowers
//         - these k flowers must be adjacent flowers
//         - we are given what day a flower blooms, only then we can pick
//         - we want to find what day ( smallest possible day ) we can make m bouquets
//     */

//     // we can start from the last possible day and keep going back from that day
//     // we know on the last day, everything is usable, therefore for sure we can make m bouquets
//     // so lets start from that day, and save the day that worked as our answer
//     // then day--, until a day that no longer works
//     // time = o(maxDay * n)
//     day := -1
//     // time = o(n)
//     for i := 0; i < len(bloomDay); i++ {
//         day = max(day, bloomDay[i])
//     }
//     res := -1
//     // time = o(maxDay * n)
//     for day >= 1 {
//         flowers := 0
//         b := 0
//         for i := 0; i < len(bloomDay); i++ {
//             if bloomDay[i] <= day {
//                 flowers++
//                 if flowers == k {b++; flowers = 0}
//             } else {
//                 flowers = 0
//             }
//         }
//         if b >= m {res = day; day--; continue}
//         break
//     }
//     return res
    
// }