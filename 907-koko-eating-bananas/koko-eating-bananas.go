func minEatingSpeed(piles []int, h int) int {
    if h == 0 {return -1}
    if piles == nil || len(piles) == 0 {return 0}
    total := 0
    // o(n)
    for i := 0; i < len(piles); i++ {
        total += piles[i]
    }
    left := 1
    right := total
    ans := 0
    for left <= right {
        mid := left + (right - left)/2
        hoursTaken := calculateHours(piles, mid)
        if hoursTaken > h {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }

    return ans
}



// func minEatingSpeed(piles []int, h int) int {
//     if h == 0 {return -1}
//     if piles == nil || len(piles) == 0 {return 0}
//     total := 0
//     // o(n)
//     for i := 0; i < len(piles); i++ {
//         total += piles[i]
//     }
//     // o(sumOfAllBananas)
//     for k := 1; k <= total; k++ {
//         // o(n)
//         nh := calculateHours(piles, k)
//         if nh <= h {
//             return k
//         }
//     }

//     // o(n) + o(sumOfAllBananas * n)
//     return -1
// }


func calculateHours(piles []int, k int) int {
    hours := 0
    for i := 0; i < len(piles); i++ {
        hours += int(math.Ceil(float64(piles[i])/float64(k)))
    }
    return hours
}