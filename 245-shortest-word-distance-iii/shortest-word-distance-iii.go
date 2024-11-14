func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    idxs := map[string][]int{}
    for idx, word := range wordsDict {
        idxs[word] = append(idxs[word], idx)
    }
    same := word1 == word2
    w1 := 0; w2 := 0
    if same {w2++}
    dist := math.MaxInt64
    for w1 < len(idxs[word1]) && w2 < len(idxs[word2]) {
        w1Idx := idxs[word1][w1]
        w2Idx := idxs[word2][w2]
        dist = min(dist, abs(w2Idx-w1Idx))
        if same {w1++; w2++} else {
            if w1Idx < w2Idx {w1++} else {w2++}
        }
    }
    return dist
}

func abs(x int) int {
    if x < 0 {return x * -1} 
    return x
}