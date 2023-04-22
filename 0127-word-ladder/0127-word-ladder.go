func ladderLength(beginWord string, endWord string, wordList []string) int {
    adjList := map[string][]string{}
    hasEndWord := false
    wordList = append(wordList, beginWord)
    for i := 0; i < len(wordList); i++ {
        parent := wordList[i]
        if parent == endWord {hasEndWord=true}
        for j := 0; j < len(wordList); j++ {
            if i == j {continue}
            child := wordList[j]
            countDiffChars := 0
            p := 0
            c := 0
            for p < len(parent) && c < len(child) {
                if parent[p] != child[c] {countDiffChars++}
                p++
                c++
            }
            if countDiffChars == 1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }
    if !hasEndWord {return 0}
    
    q := []string{beginWord}
    level := 1
    visited := map[string]struct{}{beginWord: struct{}{}}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return level}

            for _, nei := range adjList[dq] {
                if _, ok := visited[nei]; ok {continue}
                q = append(q, nei)
                visited[nei] = struct{}{}
            }
            qSize--
        }
        level++
    }
    return 0
}