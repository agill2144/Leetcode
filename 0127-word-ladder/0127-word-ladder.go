func ladderLength(beginWord string, endWord string, wordList []string) int {
    hasEnd := false    
    /*
        there could be dupes in wordList
        - we could toss wordList into a set before creating the graph
        the begin word may / may not be part of wordList
        - putting wordList in a set will help eliminate the same duplicate problem
    */ 
    wordSet := map[string]struct{}{beginWord: struct{}{}}
    for i := 0; i < len(wordList); i++ {
        if wordList[i] == endWord {hasEnd = true}
        wordSet[wordList[i]] = struct{}{}
    }    
    // if endWord does not exist, return 0, there is no path to endWord
    if !hasEnd {return 0}
    
    adjList := map[string][]string{}
    for parent, _ := range wordSet {
        for child, _ := range wordSet {
            if parent == child {continue} // skip adding same edges to the same node
            p,c := 0,0
            diff := 0
            for p < len(parent) && c < len(child) {
                if parent[p] != child[c] {
                    diff++
                }
                p++
                c++
            }
            if diff == 1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }
    // if beginWord has no adj nodes, there is no path to traverse, return 0
    if len(adjList[beginWord]) == 0 {return 0}
    
    visited := map[string]struct{}{beginWord: struct{}{}}
    q := []string{beginWord}
    level := 1
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return level}
            for _, nei := range adjList[dq] {
                if _, ok := visited[nei]; !ok {
                    q = append(q, nei)
                    visited[nei] = struct{}{}
                }
            }
            qSize--
        }
        level++
    }
    return 0
}