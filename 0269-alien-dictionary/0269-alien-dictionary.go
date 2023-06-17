func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            char := word[i]
            adjList[char] = []byte{}
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
            p++
            c++
        }
        
        if p < len(parent) && c == len(child) {return ""}
    }

    st := []byte{}
    visited := make([]bool, 26)
    var dfs func(node byte, path []bool) bool
    dfs = func(node byte, path []bool) bool {
        // base
        nodeIdx := node-'a'
        if path[nodeIdx] {return false}
        if visited[nodeIdx] {return true}
        
        // logic
        visited[nodeIdx] = true
        path[nodeIdx] = true
        for _, nei := range adjList[node] {
            if !dfs(nei, path) {return false}
        }
        path[nodeIdx] = false
        st = append(st, node)
        return true
    }

    p := make([]bool, 26)
    for i := 'a'; i <= 'z'; i++ {
        char := byte(i)
        charIdx := char-'a'
        if _, ok := adjList[char]; ok {
            if !visited[charIdx] {
                if !dfs(char, p) {return ""}
            }
        }
    }
    out := new(strings.Builder)
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        out.WriteByte(top)
    }
    return out.String()
}