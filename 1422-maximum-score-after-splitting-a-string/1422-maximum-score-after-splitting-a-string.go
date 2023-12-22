func maxScore(s string) int {
    if len(s) == 0 {return 0}
    ones := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '1' {ones++}
    }
    maxScore := math.MinInt64
    leftZeros := 0
    leftOnes := 0
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' {
            leftZeros++
        } else if s[i] == '1' {
            leftOnes++
        }
        score := leftZeros + (ones-leftOnes)
        if score > maxScore {
            maxScore = score
        }
    }
    return maxScore
}