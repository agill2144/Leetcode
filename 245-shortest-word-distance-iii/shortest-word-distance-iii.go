func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    p1 := -1; p2 := -1
    same := word1 == word2
    dist := len(wordsDict)
    for i := 0; i < len(wordsDict); i++ {
        if wordsDict[i] == word1 {
            if same {
                // if word1 == word2
                // promote current p1 value to p2
                // and p1 becomes i
                // word1, word2 = makes
                // when "makes" shows up the first time
                // p2 = -1; p1 = $firsIdx
                // when "makes" shows up the second time
                // p2 = $firstIdx; p1 = second idx
                // when "makes" shows up the third time
                // p2 = $secondIdx; p1 = third idx
                p2 = p1 
                p1 = i
            } else {
                p1 = i
            }          
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