func numSimilarGroups(strs []string) int {
    
    adjList := map[string][]string{}
    
    for i := 0; i < len(strs); i++ {
        parent := strs[i]
        for j := 0; j < len(strs); j++ {
            if i == j {continue}
            child := strs[j]
            
            p := 0
            c := 0
            diff := 0
            for p < len(parent) && c < len(child) {
                pChar := parent[p]
                cChar := child[c]
                if pChar != cChar {
                    diff++
                }
                p++; c++
            }
            if diff == 0 || diff == 2 {
                adjList[parent] = append(adjList[parent], child)
                adjList[child] = append(adjList[child], parent)
            }
        }
    }
    
    visited := map[string]struct{}{}
    var dfs func(node string)
    dfs = func(node string) {
        // base
        if _, ok := visited[node]; ok {return}
        
        // logic
        visited[node] = struct{}{}
        for _, nei := range adjList[node] {
            dfs(nei)
        }
    }

    count := 0
    for i := 0; i < len(strs); i++ {
        if _, ok := visited[strs[i]]; !ok {
            count++
            dfs(strs[i])
        }
    }
    return count
}