func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    i1 := -1
    i2 := -1
    min := math.MaxInt64

    if word1 != word2 {
        for i := 0; i < len(wordsDict); i++ {
            if wordsDict[i] == word1 {i1 = i}
            if wordsDict[i] == word2 {i2 = i}
            if i1 != -1 && i2 != -1 {
                if abs(i1-i2) < min {min = abs(i1-i2)}
            }
            if min == 1 { break }
        }
    } else { // word1 and word2 are the same
        for i := 0; i < len(wordsDict); i++ {
            if wordsDict[i] == word1 {
                if i1 < i2 {
                    i1 = i
                } else {
                    i2 = i
                }
                if i1 != -1 && i2 != -1 {
                    if abs(i1-i2) < min {min = abs(i1-i2)}
                }
            }
            if min == 1 {
                break
            }
        }
    }
    return min
}


func abs(n int) int {
    if n < 0 {
        return n*-1
    }
    return n
}