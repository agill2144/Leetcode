// wrapping / circular winow : flowing from end and wrapping around from start
func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    i := n-k
    left := n-k
    maxPoints := 0
    sum := 0
    // <= n because we could not pick any cards from the back, and pick all from the front
    // when when left == n; handle idx out of bound cases
    for left <= n {
        sum += cardPoints[i % n]
        if i-left+1 == k {
            maxPoints = max(maxPoints, sum)
            sum -= cardPoints[left%n]
            left++
        }
        i++        
    }
    return maxPoints
}

// inverted window
// func maxScore(cardPoints []int, k int) int {
//     total := 0
//     n := len(cardPoints)
//     for i := 0; i < n; i++ {total += cardPoints[i]}
//     if k >= n {return total }

//     invertedWindowSize := n-k
//     rSum := 0
//     left := 0
//     maxPoints := 0
//     for i := 0; i < n; i++ {
//         rSum += cardPoints[i]
//         if i-left+1 == invertedWindowSize {
//             maxPoints = max(maxPoints, total - rSum)
//             rSum -= cardPoints[left]
//             left++
//         }
//     }
//     return maxPoints
// }