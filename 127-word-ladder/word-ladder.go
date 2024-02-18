func ladderLength(beginWord string, endWord string, wordList []string) int {
    if beginWord == endWord {return 0}
    adjList := map[string][]string{}
    wordList = append(wordList, beginWord)
    for i := 0; i < len(wordList); i++ {
        parent := wordList[i]
        for j := 0; j < len(wordList); j++ {
            if i == j {continue}
            child := wordList[j]
            count := 0
            p, c := 0, 0
            for p < len(parent) && c < len(child) {
                if parent[p] != child[c] {
                    count++
                }
                p++
                c++
            }
            if count <= 1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }

    q := []string{beginWord}
    visited := map[string]struct{}{}
    visited[beginWord] = struct{}{}

    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {
                return level+1
            }
            for _, adjNei := range adjList[dq] {
                if _, ok := visited[adjNei]; !ok {
                    visited[adjNei] = struct{}{}
                    q = append(q, adjNei)
                }
            }
            qSize--
        }
        level++
    }
    return 0
}