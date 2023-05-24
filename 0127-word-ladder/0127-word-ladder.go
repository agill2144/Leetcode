func ladderLength(beginWord string, endWord string, wordList []string) int {
    adjList := map[string][]string{}
    
    wordList = append(wordList, beginWord)
    foundEnd := false
    for i := 0; i < len(wordList); i++ {
        parent := wordList[i]
        if parent == endWord {foundEnd = true}
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
    if !foundEnd {return 0}
    q := []string{beginWord}
    visited := map[string]struct{}{}
    visited[beginWord] = struct{}{}
    
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