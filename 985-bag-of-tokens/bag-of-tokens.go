func bagOfTokensScore(tokens []int, power int) int {
    if power <= 0 {return 0}
    // play face up -> score++
    // play face down -> score--; power+=tokens[i]\
    sort.Ints(tokens)
    left := 0
    right := len(tokens)-1
    score := 0
    maxScore := 0
    for left <= right {
        if power >= tokens[left] {
            score++
            power -= tokens[left]
            left++
        } else if left <= right && score > 0 {
            score--
            power += tokens[right]
            right--
        } else {
            return maxScore
        }
        maxScore = max(maxScore, score)
    }
    return maxScore
}

/*
    tokens = [55,71,82]
             l      r
    l = 0
    r = 2
    power = 54
*/