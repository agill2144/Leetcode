func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        word := words[i]
        for k := 0; k < len(word); k++ {
            adjList[word[k]] = []byte{}
        }
    }
    indegrees := make([]int, 26)
    for i := 0; i < len(words)-1; i++ {
        parent := words[i]
        child  := words[i+1]
        p, c := 0,0
        for p < len(parent) && c < len(child) {
            if parent[p] != child[c] {
                adjList[parent[p]] = append(adjList[parent[p]], child[c])
                indegrees[child[c]-'a']++
                break
            }
            p++
            c++
        }
        if p < len(parent) && c == len(child) {return ""}
    }
    
    q := []byte{}
    visited := make([]bool, 26)
    for k, _ := range adjList {
        if indegrees[k-'a'] == 0 {
            q = append(q, k)
            visited[k-'a'] = true
        }
    }
    
    out := new(strings.Builder)
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        out.WriteByte(dq)
        for _, nei := range adjList[dq] {
            indegrees[nei-'a']--
            if indegrees[nei-'a'] == 0 && !visited[nei-'a'] {
                q = append(q, nei)
                visited[nei-'a'] = true
            }
        }
    }
    
    for i := 0; i < len(indegrees); i++ {
        if indegrees[i] > 0 {return ""}
    }
    return out.String()
}