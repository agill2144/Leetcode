func maxScore(s string) int {
    ones := 0
    for i := 0; i < len(s); i++ {if s[i]=='1'{ones++}}
    score := 0
    countOnes, countZeros := 0, 0
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '1' {
            countOnes++
        } else {
            countZeros++
        }
        onesOnRight := ones-countOnes
        zerosOnLeft := countZeros
        score = max(score, zerosOnLeft+onesOnRight)
    }
    return score

}