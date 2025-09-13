// maximize the min 
// min = mid
// maximize will be used to take binary search to the bigger side if mid worked
func maxLength(ribbons []int, k int) int {
    left := 1
    right := slices.Max(ribbons)
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for j := 0; j < len(ribbons); j++ {
            count += int(math.Floor(float64(ribbons[j])/float64(mid)))
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
//     end := slices.Max(ribbons)
//     ans := 0
//     for i := start; i <= end; i++ {
//         count := 0
//         for j := 0; j < len(ribbons); j++ {
//             count += int(math.Floor(float64(ribbons[j])/float64(i)))
//         }
//         if count >= k {
//             ans = i
//         } else{
//             break
//         }
//     }
//     return ans
// }