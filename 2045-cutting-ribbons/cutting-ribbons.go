func maxLength(ribbons []int, k int) int {
    left := 1
    right := math.MaxInt64
    for i := 0; i < len(ribbons); i++ {
        right = max(right, ribbons[i])
    }
    ans := 0
    for left <= right {
        mid := left + (right-left)/2

        count := 0
        for j := 0; j < len(ribbons); j++ {
            count += (ribbons[j]/mid)
        }
        // when does mid work ?
        // if we were able to collect atleast k ribbons
        if count >= k {
            ans = mid
            // but keep searching right, since we want max such ans
            left = mid+1
        } else {
            // our desired size is too big, search left and our desired size of a ribbon smaller
            right = mid-1
        }
    }
    return ans
}

// tle
// func maxLength(ribbons []int, k int) int {
//     start := 1
//     end := math.MinInt64
//     for i := 0; i < len(ribbons); i++ {
//         end = max(end, ribbons[i])
//     }

//     for i := end; i >= start; i-- {
//         size := i
//         count := 0
//         // fmt.Println("desired size: ", size)
//         for j := 0; j < len(ribbons); j++ {
//             // fmt.Println("we could collect ", ribbons[j]/size, " ribbons from ", ribbons[j])
//             count += (ribbons[j]/size)
//         }
//         // when does size work
//         // when we have collected k ribbons
//         if count >= k {
//             return size
//         }
//     }
//     return 0
// }