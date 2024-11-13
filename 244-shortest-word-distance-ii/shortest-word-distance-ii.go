type WordDistance struct {
    idxs map[string][]int    
}


func Constructor(wordsDict []string) WordDistance {
    idxs := map[string][]int{}
    for i := 0; i < len(wordsDict); i++ {
        idxs[wordsDict[i]] = append(idxs[wordsDict[i]],i)
    }
    return WordDistance{idxs}    
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    p1 := 0; p2 := 0
    dist := math.MaxInt64
    for p1 < len(this.idxs[word1]) && p2 < len(this.idxs[word2]){
        idx1 := this.idxs[word1][p1]
        idx2 := this.idxs[word2][p2]
        dist = min(dist, abs(idx2-idx1))
        if idx1 < idx2 {p1++} else {p2++}
    }
    return dist
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */