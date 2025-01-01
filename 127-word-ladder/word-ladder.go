func ladderLength(beginWord string, endWord string, words []string) int {
    adjList := map[string][]string{}
    words = append(words, beginWord)
    foundEnd := false
    for i := 0; i < len(words); i++ {
        if words[i] == endWord {foundEnd = true}
        parent := words[i]
        for j := 0; j < len(words); j++ {
            if i == j {continue}
            child := words[j]
            p, c := 0,0
            diff := 0
            for p < len(parent) && c < len(child) {
                if parent[p] != child[c] {diff++}
                p++
                c++ 
            }
            if diff <= 1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }
    if !foundEnd {return 0}
    level := 1
    q := []string{beginWord}
    visited := map[string]bool{beginWord:true}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return level}
            for _, nei := range adjList[dq] {
                if !visited[nei] {
                    if nei == endWord {return level+1}
                    visited[nei] = true
                    q = append(q, nei)
                }
            }
            qSize--
        }
        level++
    }    
    return 0
}