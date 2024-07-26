func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    total := 0
    for i := 0; i < n; i++ {
        total += cardPoints[i]
    }
    if k >= n {return total}
    left := 0 
    size := n-k
    rSum := 0
    maxPoints := 0

    for i := 0; i < n; i++ {
        rSum += cardPoints[i]
        if i-left+1 == size {
            ansSum := total-rSum
            maxPoints = max(maxPoints, ansSum)
            rSum -= cardPoints[left]
            left++
        }
    }
    return maxPoints
}