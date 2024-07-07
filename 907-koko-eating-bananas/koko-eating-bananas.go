func minEatingSpeed(piles []int, h int) int {
    start := 1
    end := math.MinInt64
    for i := 0; i < len(piles); i++ {
        end = max(end, piles[i])
    }

    left := start
    right := end
    ans := -1
    for left <= right {
        mid := left + (right-left)/2

        hours := 0
        for i := 0; i < len(piles); i++ {
            // piles[i] > k; 10 and k = 1; 10 hours to finish this pile
            // k = 2; piles[i] = 10; 5 hours
            // k = 2; piles[i] = 11; 6 hours
            p := float64(piles[i])
            speed := float64(mid)
            timeTaken := int(math.Ceil(p/speed))
            hours += timeTaken
        }
        
        // when does mid not work?
        // when time taken to finish > h
        if hours > h {
            // eating speed is slow, increase it
            left = mid+1
        } else {
            // mid worked, save this as potential ans
            // because we want smallest such eating speed
            // therefore save and continue searching left
            ans = mid
            right = mid-1
        }
    }
    return ans
}
// time = o(p) + o(maxPileNum * p)
// space = o(1)
// func minEatingSpeed(piles []int, h int) int {
//     start := 1
//     end := math.MinInt64
//     for i := 0; i < len(piles); i++ {
//         end = max(end, piles[i])
//     }

//     for k := start; k <= end; k++ {
//         hours := 0
//         for i := 0; i < len(piles); i++ {
//             // piles[i] > k; 10 and k = 1; 10 hours to finish this pile
//             // k = 2; piles[i] = 10; 5 hours
//             // k = 2; piles[i] = 11; 6 hours
//             p := float64(piles[i])
//             speed := float64(k)
//             timeTaken := int(math.Ceil(p/speed))
//             hours += timeTaken
//         }
//         if hours <= h {return k}
//     }
//     return -1
// }