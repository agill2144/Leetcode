type WordDistance struct {
    idxs map[string][]int
}


func Constructor(wordsDict []string) WordDistance {
    idxs := map[string][]int{}
    for i := 0; i < len(wordsDict); i++ {idxs[wordsDict[i]] = append(idxs[wordsDict[i]], i)}
    return WordDistance{idxs}    
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    w1, w2 := 0, 0
    dist := math.MaxInt64
    for w1 < len(this.idxs[word1]) && w2 < len(this.idxs[word2]) {
        w1Idx := this.idxs[word1][w1]
        w2Idx := this.idxs[word2][w2]
        dist = min(dist, abs(w2Idx-w1Idx))
        if w1Idx < w2Idx {w1++} else {w2++}
    }
    return dist
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}


/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */