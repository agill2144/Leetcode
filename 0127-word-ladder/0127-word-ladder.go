func ladderLength(beginWord string, endWord string, wordList []string) int {
    hasEnd := false    
    /*
        there could be dupes in wordList
        - we could toss wordList into a set before creating the graph
        the begin word may / may not be part of wordList
        - putting wordList in a set will help eliminate the same duplicate problem
        
        for now, we are going to continue using the wordList as is with beginWord added
    */ 
    adjList := map[string][]string{}
    wordList = append(wordList, beginWord)
    for i := 0; i < len(wordList); i++ {
        parent := wordList[i]
        if parent == endWord {
            hasEnd = true
            continue // no need to create adjNodes for endWord, we are not going past endWord
        }
        for j := 0; j < len(wordList); j++ {
            if i == j {continue}
            child := wordList[j]
            p, c := 0, 0
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
    // if endWord is not in wordList, there is no path to reach dest, return 0
    // if beginWord has no adj nodes, there is no path to traverse, return 0
    if !hasEnd || len(adjList[beginWord]) == 0 {return 0}
    
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