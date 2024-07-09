// optimization on earlier binary search;
// we dont have to start on day 1;
// what if all the values were [6,6,6,6,6,7]
// then there is no way we would be able to make any bouquets until day 6
// i.e smallest/earliest day
// therefore update left to smallest/earliest day
// time: o( (log$maxDay-$minDay+1) * n)
// sapce = o(1)
func minDays(bloomDay []int, m int, k int) int {
    n := len(bloomDay)

    left := math.MaxInt64 // earliest possible day we can make m bouquets
    right := math.MinInt64 // last possible day we can make m bouquets
    for i := 0; i < n; i++ { 
        left = min(left, bloomDay[i])
        right = max(right, bloomDay[i])
    }

    ans := -1
    for left <= right {
        // day to evaluate
        mid := left + (right-left)/2
        bCount := 0
        fCount := 0
        for i := 0; i < n; i++ {
            if bloomDay[i] <= mid {
                fCount++
                if fCount == k {
                    bCount++
                    fCount = 0
                }
            } else {
                fCount = 0
            }
        }

        // when does mid not work?
        // mid does not work if bCount < m
        if bCount < m {
            left = mid+1
        } else {
            // mid works when we were able to make atleast m bouquets ( if not more )
            // save this as potential answer, and continue searching left because we can smallest such day
            ans = mid
            right = mid-1
        }
    }
    return ans
}


// time: o(logMaxDay * n)
// space = o(1)
// func minDays(bloomDay []int, m int, k int) int {
//     n := len(bloomDay)

//     left := 1 // earliest possible day we can make m bouquets
//     right := 0 // last possible day we can make m bouquets
//     for i := 0; i < n; i++ { right = max(right, bloomDay[i]) }

//     ans := -1
//     for left <= right {
//         // day to evaluate
//         mid := left + (right-left)/2
//         bCount := 0
//         fCount := 0
//         for i := 0; i < n; i++ {
//             if bloomDay[i] <= mid {
//                 fCount++
//                 if fCount == k {
//                     bCount++
//                     fCount = 0
//                 }
//             } else {
//                 fCount = 0
//             }
//         }

//         // when does mid not work?
//         // mid does not work if bCount < m
//         if bCount < m {
//             left = mid+1
//         } else {
//             // mid works when we were able to make atleast m bouquets ( if not more )
//             // save this as potential answer, and continue searching left because we can smallest such day
//             ans = mid
//             right = mid-1
//         }
//     }
//     return ans
// }

// time = o(n) + o(maxDay * n)
// space = o(1)
// func minDays(bloomDay []int, m int, k int) int {
//     n := len(bloomDay)
//     day := 1
//     lastPossibleDay := 0
//     for i := 0; i < n; i++ {
//         lastPossibleDay = max(lastPossibleDay, bloomDay[i])
//     }
//     for day <= lastPossibleDay {
//         bCount := 0
//         fCount := 0
//         for i := 0; i < n; i++ {
//             if bloomDay[i] <= day {
//                 fCount++
//                 if fCount == k {
//                     bCount++
//                     fCount = 0
//                     if bCount == m {return day}
//                 }
//             } else {
//                 fCount = 0
//             }
//         }
//         day++
//     }
//     return -1
// }