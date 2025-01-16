func ladderLength(beginWord string, endWord string, wordList []string) int {
    words := map[string]bool{}
    foundEnd := false
    //o(n) where n is len words
    // k = avg len of each word
    for i := 0; i < len(wordList); i++ {
        if wordList[i] == endWord{foundEnd = true}
        words[wordList[i]] = true
    }
    
    if !foundEnd {return 0}
    
    q := []string{beginWord}
    levels := 1
    visited := map[string]bool{beginWord:true}
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
                    child := dq[:i] + string(j+'a') + dq[i+1:]
                    if !visited[child] && words[child] {
                        q = append(q, child)
                        visited[child] = true
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