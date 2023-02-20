func ladderLength(beginWord string, endWord string, wordList []string) int {
    adjList := map[string][]string{}
    wordList = append(wordList, beginWord)
    for i := 0; i < len(wordList); i++ {
        parent := wordList[i]
        for j := 0; j < len(wordList); j++ {
            if j == i {continue}
            child := wordList[j]
            p := 0
            c := 0
            count := 0
            for p < len(parent) {
                if parent[p] == child[c] {count++}
                p++
                c++
            }
            if count == len(parent)-1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }
    visitedWord := map[string]struct{}{beginWord: struct{}{}}
    q := []string{beginWord}
    level := 1
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]; q = q[1:]
            if dq == endWord { return level } 
            for _, child := range adjList[dq] {
                if _, ok := visitedWord[child]; !ok {
                    q = append(q, child)
                    visitedWord[child] = struct{}{}
                }
            }
            qSize--
        }
        level++
    }
    return 0
}