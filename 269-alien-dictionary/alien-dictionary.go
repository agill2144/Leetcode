func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        for k := 0; k < len(words[i]); k++ {
            adjList[words[i][k]] = []byte{}
        }
    }

    for i := 0; i < len(words)-1; i++ {
        parent := words[i]
        child := words[i+1]
        p, c := 0, 0
        for p < len(parent) && c < len(child) {
            if parent[p] != child[c] {
                adjList[parent[p]] = append(adjList[parent[p]], child[c])
                break
            }
            p++; c++
        }
        if c == len(child) && p < len(parent) {return ""}
    }
    st := []byte{}
    visited := make([]bool, 26)
    var dfs func(node byte, path []bool) bool
    dfs = func(node byte, path []bool) bool {
        // base
        if path[int(node-'a')] {return false}
        if visited[int(node-'a')] {return true}

        // logic
        visited[int(node-'a')] = true
        path[int(node-'a')] = true
        for _, nei := range adjList[node] {
            if !dfs(nei, path) {return false}
        }
        st = append(st, node)
        path[int(node-'a')] = false
        return true
    }
    p := make([]bool, 26)
    for k, _ := range adjList {
        if !dfs(k, p){return ""}
    }
    res := &strings.Builder{}
    for i := len(st)-1; i >= 0; i-- {
        res.WriteByte(st[i])
    }
    return res.String()
}