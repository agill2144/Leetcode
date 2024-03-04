func bagOfTokensScore(tokens []int, power int) int {
    // when we have no power
    if power <= 0 {return 0}
    
    // sort the tokens so that we can play greedy
    sort.Ints(tokens)
    

    // how are we going to play greedily ?
    // well we want to maximize our score
    // scores goes up if we can play a token
    // we can play a token if we power >= token value
    // what if we do not have enough power ?
    // ( thats why we sorted )
    // we want to increase our power as much as possible when
    // we do not have enough to play
    // therefore pick the highest / max token avail
    // when we play a token face-down ( because we didnt have enough power )
    // our power goes by token-value and scores goes down by 1
    // therefore we will check first whether we have power
    // if yes, play the token and increment score
    // if no, play the token face-down and add the token value to power
    // if none of the above 2 works ( even after adding more power )
    // this is the best we can do and we can exit
    // time = o(nlogn) + o(n)
    // space = o(1)
    runningScore := 0
    score := 0
    left, right := 0, len(tokens)-1
    for left <= right {
        if power >= tokens[left] {
            runningScore++
            power -= tokens[left]
            left++
        } else if runningScore >= 1 {
            power += tokens[right]
            right--
            runningScore--
        } else {
            break
        }
        score = max(score, runningScore)
    }
    return score
}