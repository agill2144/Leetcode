func ladderLength(beginWord string, endWord string, wordList []string) int {
    foundEnd := false
    wordSet := map[string]struct{}{}
    for i := 0; i < len(wordList); i++ {
        word := wordList[i]
        if word == endWord {foundEnd = true}
        wordSet[word] = struct{}{}
    }
    if !foundEnd {return 0}
    
    q := []string{beginWord}
    levels := 1
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return levels}
            for i := 0; i < len(dq); i++ {
                for j := 0; j < 26; j++ {
                    char := byte(j+'a')
                    nei := dq[:i] + string(char) + dq[i+1:]
                    if _, ok := wordSet[nei]; ok {
                        if nei == endWord {return levels+1}
                        delete(wordSet, nei)
                        q = append(q, nei)
                    }
                }
            }
            qSize--
        }
        levels++
    }
    return 0
}