func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    idxs := map[string][]int{}
    for i := 0; i < len(wordsDict); i++ {
        idxs[wordsDict[i]] = append(idxs[wordsDict[i]],i)
    }
    p1 := 0
    p2 := 0
    dist := len(wordsDict)
    for p1 < len(idxs[word1]) && p2 < len(idxs[word2]) {
        idx1 := idxs[word1][p1]
        idx2 := idxs[word2][p2]
        dist = min(dist, abs(idx2-idx1))
        if idx1 < idx2 {
            p1++
        } else {
            p2++
        }
    }
    return dist
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}
