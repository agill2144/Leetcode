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
        if len(parent) > len(child) && strings.HasPrefix(parent, child) {return ""}
        p := 0
        c := 0
        for p < len(parent) && c < len(child) {
            if parent[p] != child[c] {
                adjList[parent[p]] = append(adjList[parent[p]], child[c])
                indegrees[child[c]-'a']++
                break
            }
            p++; c++
        }
        // if p < len(parent) && c >= len(child) {return ""}
    }
    
    q := []byte{}
    for key, _ := range adjList {
        if indegrees[key-'a'] == 0 {q = append(q, key)}
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
    
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] > 0 {return ""}
    }
    return out.String()
}