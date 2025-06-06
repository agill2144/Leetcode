func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    maxPoints := 0
    rSum := 0
    left := n-k
    i := n-k
    for left <= n {
        rSum += cardPoints[i%n]
        if i-left+1 == k {
            maxPoints = max(maxPoints, rSum)
            rSum -= cardPoints[left%n]
            left++
        }
        i++
    }
    return maxPoints
}
// 2 pass inverted window
// func maxScore(cardPoints []int, k int) int {
//     n := len(cardPoints)
//     total := 0
//     for i := 0; i < n; i++ {total += cardPoints[i]}
//     if k == n {return total}
//     left := 0
//     winSum := 0
//     maxPoints := 0
//     for i := 0; i < n; i++ {
//         winSum += cardPoints[i]
//         if i-left+1 == n-k {
//             maxPoints = max(maxPoints, total-winSum)
//             winSum -= cardPoints[left]
//             left++
//         }
//     }
//     return maxPoints
// }