type WordDistance struct {
    wordsDict []string    
}


func Constructor(wordsDict []string) WordDistance {
    return WordDistance{wordsDict}    
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    w1, w2 := -1, -1
    dist := len(this.wordsDict)
    for idx, word := range this.wordsDict {
        if word == word1 {w1 = idx}
        if word == word2 {w2 = idx}
        if w1 != -1 && w2 != -1 {
            dist = min(dist, abs(w2-w1))
            if dist == 1 {break}
        }
    }
    return dist    
}


func abs(x int) int {
    if x < 0 {return x * -1} 
    return x
}
/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */