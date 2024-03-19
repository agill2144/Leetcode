func shipWithinDays(weights []int, days int) int {
    left := math.MinInt64
    right := 0
    for i := 0; i < len(weights); i++ {
        if weights[i] > left {left = weights[i]}
        right += weights[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        // if atMax, we can take $mid weight per day
        // how many days would it take to take all packages

        daysCount := 1
        rSum := 0
        for i := 0; i < len(weights); i++ {
            rSum += weights[i]
            if rSum > mid {
                daysCount++
                rSum = weights[i]
            }
        }

        // did mid work?
        // mid works if it took less or equal to days allowed
        // to ship all packages
        if daysCount <= days {
            // save this as potential ans
            ans = mid
            // look left, since we are searching for smallest such ans
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

// func shipWithinDays(weights []int, days int) int {
//     start := math.MinInt64
//     end := 0
//     for i := 0; i < len(weights); i++ {
//         if weights[i] > start {start = weights[i]}
//         end += weights[i]
//     }

//     for i := start ; i <= end; i++ {
//         capacity := i // atMax capacity, cannot exceed

//         // how many days would it take if atMax our capacity is $capacity
//         totalDays := 1
//         rSum := 0
//         for j := 0; j < len(weights); j++ {
//             rSum += weights[j]
//             if rSum > capacity {
//                 totalDays++
//                 rSum = weights[j]
//             }
//         }
//         if totalDays <= days {
//             return capacity
//         }
//     }
//     return -1
// }