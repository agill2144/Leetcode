func maxScore(cardPoints []int, k int) int {
    maxPoints := 0
    n := len(cardPoints)
    left := n-k
    i := left
    points := 0
    for left <= n {
        points += cardPoints[i%n]
        if i-left+1 == k {
            maxPoints = max(maxPoints, points)
            points -= cardPoints[left%n]
            left++
        }
        i++
    }
    return maxPoints
}