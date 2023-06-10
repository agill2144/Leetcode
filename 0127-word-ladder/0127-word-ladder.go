func ladderLength(beginWord string, endWord string, wordList []string) int {
    adjList := map[string][]string{}
    foundEnd := false
    wordList = append(wordList, beginWord)
    for i := 0; i < len(wordList); i++ {
        parent := wordList[i]
        if parent == endWord {foundEnd = true}
        for j := 0; j < len(wordList); j++ {
            if j == i {continue}
            child := wordList[j]
            diff := 0
            p, c := 0,0
            for p < len(parent) && c < len(child) {
                if parent[p] != child[c] {diff++}
                p++
                c++
            }
            if diff == 1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }
    if !foundEnd {return 0}
    
    visited := map[string]struct{}{beginWord: struct{}{}}
    level := 1
    q := []string{beginWord}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return level}
            for _, nei := range adjList[dq] {
                if _, ok := visited[nei]; !ok {
                    if nei == endWord {return level+1}
                    visited[nei] = struct{}{}
                    q = append(q, nei)
                }
            }
            qSize--
        }
        level++
    }
    return 0
}