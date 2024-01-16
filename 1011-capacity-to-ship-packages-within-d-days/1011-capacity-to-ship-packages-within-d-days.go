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
        
        rSum := 0
        totalDays := 1
        atMax := mid
        for i := 0; i < len(weights); i++ {
            rSum += weights[i]
            if rSum > atMax {
                rSum = weights[i]
                totalDays++
            }
        }
        if totalDays > days {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
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

//     for i := start; i <= end; i++ {
//         atMax := i
//         rSum := 0
//         totalDays := 1
        
//         for j := 0; j < len(weights); j++ {
//             rSum += weights[j]
//             if rSum > atMax {
//                 totalDays++
//                 if rSum == atMax {rSum = 0} else {rSum = weights[j]} 
                
//             }    
//         }        
//         if totalDays <= days {return i}
//     }
//     return -1
// }