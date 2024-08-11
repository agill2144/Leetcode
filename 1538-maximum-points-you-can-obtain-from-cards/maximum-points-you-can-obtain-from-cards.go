func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    left := n-k
    rSum := 0
    i := left
    maxPoints := 0
    for left <= n {
        rSum += cardPoints[i%n]
        if i-left+1 == k {
            maxPoints = max(maxPoints, rSum)
            if left == n {break}
            rSum -= cardPoints[left]
            left++
        }
        i++
    }
    return maxPoints
}