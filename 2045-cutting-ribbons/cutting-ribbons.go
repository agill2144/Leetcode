func maxLength(ribbons []int, k int) int {
    left := 1
    right := -1
    for i := 0; i < len(ribbons); i++ { right = max(right, ribbons[i]) }
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for j := 0; j < len(ribbons); j++ {
            count += (ribbons[j]/mid)
        }
        if count >= k {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}

// func maxLength(ribbons []int, k int) int {
//     start := 1
//     end := -1
//     for i := 0; i < len(ribbons); i++ {
//         end = max(end, ribbons[i])
//     }
//     ans := 0
//     for i := start ; i <= end; i++ {
//         count := 0
//         for j := 0; j < len(ribbons); j++ {
//             count += (ribbons[j]/i)
//         }
//         if count >= k {ans = i; continue}
//         break
//     }
//     return ans
// }