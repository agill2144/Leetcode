func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        word := words[i]
        for k := 0; k < len(word); k++ {
            adjList[word[k]] = []byte{}
        }
    }
    for i := 0; i < len(words)-1; i++ {
        parent := words[i]
        child  := words[i+1]
        p, c := 0,0
        for p < len(parent) && c < len(child) {
            if parent[p] != child[c] {
                adjList[parent[p]] = append(adjList[parent[p]], child[c])
                break
            }
            p++
            c++
        }
        if p < len(parent) && c == len(child) {return ""}
    }
    
    visited := make([]bool, 26)
    st := []byte{}
    var dfs func(node byte, path []bool) bool
    dfs = func(node byte, path []bool) bool {
        // base
        if path[node-'a'] {return false}
        if visited[node-'a'] {return true}
        
        // logic
        visited[node-'a'] = true
        path[node-'a'] = true
        for _, nei := range adjList[node] {
            if !dfs(nei, path) {return false}
        }
        st = append(st, node)
        path[node-'a'] = false
        return true
    }
    
    p := make([]bool, 26)
    for k, _ := range adjList {
        if !visited[k-'a'] {
            if !dfs(k, p) {return ""}
        }
    }
    
    out := new(strings.Builder)
    for len(st) != 0 {
        out.WriteByte(st[len(st)-1])
        st = st[:len(st)-1]
    }
    return out.String()
}