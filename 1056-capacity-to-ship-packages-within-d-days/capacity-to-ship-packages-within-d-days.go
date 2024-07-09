func shipWithinDays(weights []int, days int) int {
    left := math.MinInt64
    right := 0
    for i := 0; i < len(weights); i++ {
        right += weights[i]
        left = max(left, weights[i])
    }

    // binary search on answers ( weight is what we are looking for )
    ans := -1
    for left <= right {

        mid := left + (right-left)/2
        rSum := 0
        numDays := 1
        for j := 0; j < len(weights); j++ {
            rSum += weights[j]
            if rSum > mid {
                numDays++
                rSum = weights[j]
            }
        }

        // when does mid as capacity for the ship not work ?
        // when it took more than $days to ship all packages
        if numDays > days {
            // increase the weight
            left = mid+1
        } else {
            // save this as potential ans
            ans = mid
            // keep searching left since we want smallest such ans
            right = mid-1
        }
    }
    return ans
}

// time = o((sum-max+1) * n)
// space = o(1)
// func shipWithinDays(weights []int, days int) int {
//     start := math.MinInt64
//     end := 0
//     for i := 0; i < len(weights); i++ {
//         end += weights[i]
//         start = max(start, weights[i])
//     }

//     for i := start; i <= end; i++ {

//         cap := i
//         rSum := 0
//         numDays := 1
//         for j := 0; j < len(weights); j++ {
//             rSum += weights[j]
//             if rSum > cap {
//                 numDays++
//                 rSum = weights[j]
//             }
//         }

//         if numDays <= days {return cap}
//     }
//     return -1
// }