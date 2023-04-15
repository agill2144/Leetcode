func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    indegrees := make([]int, 26)
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            adjList[words[i][j]] = []byte{}
        }
    }
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        w1, w2 := 0, 0
        for w1 < len(word1) && w2 < len(word2) && word1[w1] == word2[w2] {
            w1++
            w2++
        }
        // invalid sequence
        if w2 == len(word2) && w1 < len(word1) {return ""}
        if (w2 == len(word2) && w1 == len(word1)) || (w1 == len(word1) && w2 < len(word2)) {continue}
        adjList[word1[w1]] = append(adjList[word1[w1]], word2[w2])
        indegrees[word2[w2]-'a']++
    }
    q := []byte{}
    for char, _ := range adjList {
        if indegrees[char-'a'] == 0 {
            q = append(q, char)
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

    if len(out.String()) != len(adjList) {
        return ""
    }    
    return out.String()
}