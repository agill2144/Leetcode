func maximizeSweetness(sweetness []int, k int) int {
    // min sweetness is maximized
    start := math.MaxInt64
    end := 0
    for i := 0; i < len(sweetness); i++ {
        if sweetness[i] < start {
            start = sweetness[i]
        }
        end += sweetness[i]
    }
    ans := 0
    left := start
    right := end
    for left <= right {

        // at min, you can have this
        mid := left + (right-left)/2

        count := 0
        rSum := 0
        for j := 0; j < len(sweetness); j++ {
            rSum += sweetness[j]
            if rSum >= mid {
                count++
                rSum = 0
            }
        }
        
        if count >= k+1 {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}

// tle
// func maximizeSweetness(sweetness []int, k int) int {
//     // min sweetness is maximized
//     start := math.MaxInt64
//     end := 0
//     for i := 0; i < len(sweetness); i++ {
//         if sweetness[i] < start {
//             start = sweetness[i]
//         }
//         end += sweetness[i]
//     }
//     ans := 0    
//     for i := start; i <= end; i++ {
//         atMin := i
        
//         count := 0
//         rSum := 0
//         for j := 0; j < len(sweetness); j++ {
//             rSum += sweetness[j]
//             if rSum >= atMin {
//                 count++
//                 rSum = 0
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