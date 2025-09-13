// maximize the atMin sweetness
func maximizeSweetness(sweetness []int, k int) int {
    left := math.MaxInt64
    right := 0
    for i := 0; i < len(sweetness); i++ {
        left = min(left, sweetness[i])
        right += sweetness[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        rSum := 0
        count := 0
        for i := 0; i < len(sweetness); i++ {
            rSum += sweetness[i]
            if rSum >= mid {
                count++
                rSum = 0
                if count >= k+1 {break}
            }
        }
        // when does mid not work?
        // when we could not make atleast k+1 pieces
        if count < k+1 {
            // per chunk atMin sweetness is to high
            // hence reduce the atMin and try again in-hope to cut more pieces
            right = mid-1
        } else {
            ans = mid
            // we want max such ans, continue searching right
            left = mid+1
        }
    }
    return ans

}
// func maximizeSweetness(sweetness []int, k int) int {
//     start := math.MaxInt64
//     end := 0
//     for i := 0; i < len(sweetness); i++ {
//         start = min(start, sweetness[i])
//         end += sweetness[i]
//     }
//     ans := -1
//     for i := start; i <= end; i++ {
//         rSum := 0
//         count := 0
//         for j := 0; j < len(sweetness); j++ {
//             rSum += sweetness[j]
//             if rSum >= i {
//                 rSum = 0
//                 count++
//             }
//         }
//         if count >= k+1 {
//             ans = i
//         } else {
//             break
//         }
//     }
//     return ans
// }
