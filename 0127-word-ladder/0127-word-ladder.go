func ladderLength(beginWord string, endWord string, wordList []string) int {
    if beginWord == endWord || len(wordList) == 0 {return 0}
    wordSet := map[string]struct{}{}
    foundEnd := false
    //o(n) where n is len words
    // k = avg len of each word
    for i := 0; i < len(wordList); i++ {
        if wordList[i] == endWord{foundEnd = true}
        wordSet[wordList[i]] = struct{}{}
    }
    
    if !foundEnd {return 0}
    
    q := []string{beginWord}
    levels := 1
    // o(n)
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return levels}
            
            // o(k)  
            for i := 0; i < len(dq); i++ {
                // o(26) ; i.e o(1)
                for j := 0; j < 26; j++ {
                    char := string(byte(j+'a'))
                    // o(k); slicing takes o(k) time
                    nei := dq[:i] + char + dq[i+1:]
                    if nei == endWord {return levels+1}
                    if _, ok := wordSet[nei]; ok {
                        q = append(q, nei)
                        delete(wordSet, nei)
                    }
                }
            }
            qSize--
        }
        levels++
    }
    // total time = o(n) + o(n * k^2)
    // total space = o(n) wordSet, o(n) for queue 
    
    return 0
}