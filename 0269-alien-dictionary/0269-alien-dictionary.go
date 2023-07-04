/*
    questions to ask:
    1. guranteed sorted correctly ?
        - No: then check while creating the adjList
        - one case could be the that the 2nd string is smaller and parent is larger
            ["amrit", "amit"] - not sorted correctly
    2. Should the ouput string contain only the characters that are used to construct this sorted order or ALL characters
        - if just the chars that are used to sort, then our adjList can only contain the diff chars 
            - keep in mind, if there were no diffs detected, the output in this case will be empty ( no order could be determined )
        - if all characters of the input words should be part of output
            - we can simply add each character to adjList first and then find its childrens
            
*/
func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        for k := 0; k < len(words[i]); k++ {
            adjList[words[i][k]] = []byte{}
        }
    }
    
    // find the differences between 2 words
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
        // if all characters were same, while parent ptr is still inbound and
        // child ptr no longer has words, it means child was smaller than parent ( with same exact chars )
        // if sorted correctly, child should be placed before but in this case its not sorted correctly
        // exit early
        if p < len(parent) && c == len(child) {
            return ""
        }
    }
    
    // cycle detection algo on all chars that we care about
    // the keys of adjList
    // why not also the childs ? or just the child chars ?
    // well key is the parent, i.e key char came before child character
    // therefore key/parent has no dependencies and childs do
    st := []byte{}
    visited := make([]bool, 26)
    var dfs func(node byte, path []bool) bool
    dfs = func(node byte, path []bool) bool {
        // base
        if path[node-'a'] {return false}// cycle detected
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
        top := st[len(st)-1]
        st = st[:len(st)-1]
        out.WriteByte(top)
    }
    return out.String()
}
