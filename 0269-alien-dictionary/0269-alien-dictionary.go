func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            adjList[words[i][j]] = []byte{}
        }
    }
    indegrees := make([]int, 26)
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        w1, w2 := 0,0
        
        for w1 < len(word1) && w2 < len(word2) {
            w1Char := word1[w1]
            w2Char := word2[w2]
            if w1Char != w2Char {
                adjList[w1Char] = append(adjList[w1Char], w2Char)
                indegrees[w2Char-'a']++
                break
            }
            w1++
            w2++
        }
        if w2 == len(word2) && w1 < len(word1) {return ""}
    }
    
    q := []byte{}
    for key, _ := range adjList {
        if indegrees[key-'a'] == 0 {
            q = append(q, key)
        }
    }
    if len(q) == 0 {
        return ""
    }
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
    if len(out.String()) != len(adjList) {return ""}
    return out.String()
}