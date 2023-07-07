func maxConsecutiveAnswers(answerKey string, k int) int {
    maxWindowSize := 0
    left := 0
    tCount := 0
    fCount := 0
    for i := 0; i < len(answerKey); i++ {
        char := string(answerKey[i])
        if char == "T" {tCount++}
        if char == "F" {fCount++}
        winSize := i-left+1
        // make sure whether our window has 1 possible ans
        for winSize-tCount > k && winSize-fCount > k {
            leftChar := string(answerKey[left])
            if leftChar == "T" {tCount--}
            if leftChar == "F" {fCount--}
            left++
            winSize = i-left+1
        }
        // save the window size if better 
        if (winSize-tCount <= k || winSize-fCount <= k) && winSize > maxWindowSize {
            maxWindowSize = winSize
        } 
    }
    return maxWindowSize
}