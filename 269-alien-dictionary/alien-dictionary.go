func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        for k := 0; k < len(words[i]); k++ {
            adjList[words[i][k]] = []byte{}
        }
    }
    indegrees := make([]int, 26)
    for i := 0; i < len(words)-1; i++ {
        parent := words[i]
        child := words[i+1]
        p,c := 0, 0
        for p < len(parent) && c < len(child) {
            if parent[p] != child[c] {
                adjList[parent[p]] = append(adjList[parent[p]], child[c])
                indegrees[int(child[c]-'a')]++
                break
            }
            p++; c++
        }
        // this is not lexicographically sorted!
        if c == len(child) && p < len(parent) {
            return ""
        }
    }
    countToBeProcessed := 0
    q := []byte{}
    for i := 0; i < len(indegrees); i++ {
        char := byte('a'+i)
        if adjList[char] != nil {
            countToBeProcessed++
            if indegrees[i] == 0 {q = append(q, char)}
        }
    }
    if len(q) == 0 {return ""}
    res := new(strings.Builder)
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        countToBeProcessed--
        res.WriteByte(dq)
        for _, nei := range adjList[dq] {
            indegrees[int(nei-'a')]--
            if indegrees[int(nei-'a')] == 0 {
                q = append(q, nei)
            }
        }
    }
    if countToBeProcessed != 0 {return ""}
    return res.String()
    
}