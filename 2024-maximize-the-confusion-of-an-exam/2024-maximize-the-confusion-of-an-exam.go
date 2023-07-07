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
        for winSize-tCount > k && winSize-fCount > k {
            leftChar := string(answerKey[left])
            if leftChar == "T" {tCount--}
            if leftChar == "F" {fCount--}
            left++
            winSize = i-left+1
        }
        if (winSize-tCount <= k || winSize-fCount <= k) && winSize > maxWindowSize {
            maxWindowSize = winSize
        } 
    }
    return maxWindowSize
}