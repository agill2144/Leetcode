func maxScore(cardPoints []int, k int) int {
    total := 0
    n := len(cardPoints)
    for i := 0; i < n; i++ {total += cardPoints[i]}
    if k >= n {return total }
    
    invertedWindowSize := n-k
    rSum := 0
    left := 0
    maxPoints := 0
    for i := 0; i < n; i++ {
        rSum += cardPoints[i]
        if i-left+1 == invertedWindowSize {
            maxPoints = max(maxPoints, total - rSum)
            rSum -= cardPoints[left]
            left++
        }
    }
    return maxPoints
}