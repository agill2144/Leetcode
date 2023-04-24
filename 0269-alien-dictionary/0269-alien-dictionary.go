func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    
    for i := 0; i < len(words); i++ {
        word := words[i]
        for k := 0; k < len(word); k++ {
            adjList[word[k]] = []byte{}
        }
    }
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]        
        
        w1, w2 := 0,0
        
        for w1 < len(word1) && w2 < len(word2) {
            if word1[w1] != word2[w2] {
                adjList[word1[w1]] = append(adjList[word1[w1]], word2[w2])
                break
            }
            w1++
            w2++
        }
        
        // invalid input
        if w2 == len(word2) && w1 < len(word1) {return ""}
    }
    st := []byte{}
    visited := map[byte]struct{}{}
    var dfs func(node byte, path map[byte]struct{}) bool
    dfs = func (node byte, path map[byte]struct{}) bool {
        // base
        if _, ok := path[node]; ok {return false}
        if _, ok := visited[node]; ok {return true}
        
        // logic
        visited[node] = struct{}{}
        path[node] = struct{}{}
        for _, nei := range adjList[node] {
            if !dfs(nei, path) {return false}
        }
        st = append(st, node)
        delete(path, node)
        return true
    }
    for k, _ := range adjList {
        if _, ok := visited[k]; !ok {
            if !dfs(k, map[byte]struct{}{}) {return ""}
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