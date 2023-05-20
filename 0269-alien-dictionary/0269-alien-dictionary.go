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
        p := 0
        c := 0
        for p < len(parent) && c < len(child) {
            if parent[p] != child[c] {
                adjList[parent[p]] = append(adjList[parent[p]], child[c])
                break
            }
            p++; c++
        }
        // when words are sorted alphabetically but incorrectly sorted by length
        // i.e [abcdef, abc]
        // abc should come before abcdef, therefore this is invalid
        // and this is when our parent ptr is still inbound while child is out-of-bounds, in that case, return ""
        if p < len(parent) && c >= len(child) {return ""}
    }
    visited := make([]bool, 26)
    st := []byte{}
    
    // classic cycle detection on directed graph
    // using dfs
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
        if !visited[k-'a'] && !dfs(k, p) {return ""}
    }
    out := new(strings.Builder)
    for i := len(st)-1; i >= 0; i-- {
        out.WriteByte(st[i])
    }
    return out.String()
}

// func alienOrder(words []string) string {
//     adjList := map[byte][]byte{}
    
//     for i := 0; i < len(words); i++ {
//         for k := 0; k < len(words[i]); k++ {
//             adjList[words[i][k]] = []byte{}
//         }
//     }
    
//     indegrees := make([]int, 26)
//     for i := 0; i < len(words)-1; i++ {
//         parent := words[i]
//         child := words[i+1]
//         p := 0
//         c := 0
//         for p < len(parent) && c < len(child) {
//             if parent[p] != child[c] {
//                 adjList[parent[p]] = append(adjList[parent[p]], child[c])
//                 indegrees[child[c]-'a']++
//                 break
//             }
//             p++; c++
//         }
//         // when words are sorted alphabetically but incorrectly sorted by length
//         // i.e [abcdef, abc]
//         // abc should come before abcdef, therefore this is invalid
//         // and this is when our parent ptr is still inbound while child is out-of-bounds, in that case, return ""
//         if p < len(parent) && c >= len(child) {return ""}
//     }
    
//     q := []byte{}
//     /*
//        why not use indegrees array as is?
//        - the indegrees array is an int array
//        - some characters may not exist in the input word list but their indegrees is 0
//        - and we do not care about all characters that have 0 incoming edges
//        - we only care about characters from our input word list that has indgrees of 0
       
//        why loop over keys?
//        - because adjList has a format of {parent: child} , parent --> child
//        - which means, child will always have an incoming edge
//        - and we only want to find our start nodes that have incoming edge of 0
//        - those nodes will never be the child nodes, they can be the parent ones
//        - therefore looping over keys of adjList to find out nodes we care about
       
//     */ 
//     for key, _ := range adjList {
//         if indegrees[key-'a'] == 0 {q = append(q, key)}
//     }
//     out := new(strings.Builder)
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         out.WriteByte(dq)
//         for _, nei := range adjList[dq] {
//             indegrees[nei-'a']--
//             if indegrees[nei-'a'] == 0 {
//                 q = append(q, nei)
//             }
//         }
//     }
    
//     // check if there are any nodes that could not be processed
//     // this means there were cycles, when thats the case, we can return ""
//     for i := 0; i < len(indegrees); i++ {
//         if indegrees[i] > 0 {return ""}
//     }
    
//     return out.String()
// }