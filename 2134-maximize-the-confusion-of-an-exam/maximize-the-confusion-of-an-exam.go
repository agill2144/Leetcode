func maxConsecutiveAnswers(answerKey string, k int) int {
    ts := 0
    fs := 0
    maxCount := 0
    left := 0
    maxSize := 0
    for i := 0; i < len(answerKey); i++ {
        if answerKey[i] == 'T' {ts++}
        if answerKey[i] == 'F' {fs++}
        maxCount = max(maxCount, max(ts, fs))
        size := i-left+1
        smallerCount := size-maxCount
        if smallerCount <= k {
            maxSize = max(maxSize, i-left+1)
        } else {
            if answerKey[left] == 'T' {ts--}
            if answerKey[left] == 'F' {fs--}
            maxCount = max(maxCount, max(ts, fs))
            left++
        }
    }
    return maxSize
}