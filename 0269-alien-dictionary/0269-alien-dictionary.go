func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    
    // [ "a", "a" ]
    
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            adjList[words[i][j]] = []byte{}
        }
    }
    indegrees := make([]int, 26)
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        
        if strings.HasPrefix(word1, word2) && len(word2) < len(word1) {
            // invalid order
            return ""
        }
        
        for j := 0; j < len(word1) && j < len(word2); j++ {
            w1Char := word1[j]
            w2Char := word2[j]
            if w1Char != w2Char {
                adjList[w1Char] = append(adjList[w1Char], w2Char)
                indegrees[w2Char-'a']++
                break
            }
        }
    }
    q := []byte{}
    for key, _ := range adjList {
        if indegrees[key-'a'] == 0 {
            q = append(q, key)
        }
    }
    if len(q) == 0 {return ""}
    out := new(strings.Builder)
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        out.WriteByte(dq)
        for _, nei := range adjList[dq] {
            indegrees[nei-'a']--
            if indegrees[nei-'a'] == 0 {
                q = append(q, nei)
            }
        }
    }
    // if there are still dependents not fully resolved, that means there was a cycle.
    for i := 0; i < 26; i++ {
        if indegrees[i] != 0 {return ""}
    }
    return out.String()
}