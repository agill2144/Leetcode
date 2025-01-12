func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            adjList[words[i][j]] = []byte{}
        }
    }
    for i := 0; i < len(words)-1; i++ {
        parent := words[i]
        child := words[i+1]
        p, c := 0,0
        for p < len(parent) && c < len(child) {
            if parent[p] == child[c] {p++; c++; continue}
            adjList[parent[p]] = append(adjList[parent[p]], child[c])
            break
        }
        if c == len(child) && p < len(parent) {return ""}
    }
    
    st := []byte{}
    visited := make([]bool, 26)    
    var dfs func(curr byte, path []bool) bool
    dfs = func(curr byte, path []bool) bool {
        // base
        if path[int(curr-'a')] {return false}
        if visited[int(curr-'a')] {return true}

        // logic
        visited[int(curr-'a')] = true
        path[int(curr-'a')] = true
        for _, nei := range adjList[curr] {
            if !dfs(nei, path) {return false}
        }
        path[int(curr-'a')] = false
        st = append(st, curr)
        return true
    }
    p := make([]bool, 26)
    for k, _ := range adjList {
        if !dfs(k, p) {return ""}
    }
    res := new(strings.Builder)
    for len(st) != 0{
        top := st[len(st)-1]
        st = st[:len(st)-1]
        res.WriteByte(top)
    }
    return res.String()
}