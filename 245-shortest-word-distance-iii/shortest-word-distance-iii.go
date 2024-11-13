func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    p1 := -1; p2 := -1
    same := word1 == word2
    dist := len(wordsDict)
    for i := 0; i < len(wordsDict); i++ {
        if wordsDict[i] == word1 {
            if same {
                p2 = p1
                p1 = i
            } else {p1 = i}          
        } else if wordsDict[i] == word2 {
            p2 = i
        }
        if p1 != -1 && p2 != -1 {dist = min(dist, abs(p2-p1))}
    }
    return dist
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}