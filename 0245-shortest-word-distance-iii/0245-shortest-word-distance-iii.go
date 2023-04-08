func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    w1Idxs := []int{}
    w2Idxs := []int{}
    for i := 0; i < len(wordsDict); i++ {
        if wordsDict[i] == word1 {w1Idxs = append(w1Idxs, i)}
        if wordsDict[i] == word2 {w2Idxs = append(w2Idxs, i)}
    }
    w1Ptr := 0
    w2Ptr := 0
    minDiff := math.MaxInt64
    for w1Ptr < len(w1Idxs) && w2Ptr < len(w2Idxs) {
        w1Idx := w1Idxs[w1Ptr]
        w2Idx := w2Idxs[w2Ptr]
        currDiff := abs(w1Idx-w2Idx)
        if currDiff == 0 {
            w1Ptr++
            continue
        }
        if currDiff < minDiff {
            minDiff = currDiff
        }
        // which ptr to move
        if w2Idx < w1Idx {
            w2Ptr++
        } else {
            w1Ptr++
        }
    }
    return minDiff
}


func abs(x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}