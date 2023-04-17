func ladderLength(beginWord string, endWord string, wordList []string) int {
    adjList := map[string][]string{}
    wordList = append(wordList, beginWord)
    listHasEndWord := false
    for i := 0; i < len(wordList); i++ {
        parentWord := wordList[i]
        if parentWord == endWord {listHasEndWord = true}
        for j := 0; j < len(wordList); j++ {
            if i == j {continue}
            child := wordList[j]
            countDiffChars := 0
            p := 0
            c := 0
            for p < len(parentWord) && c < len(child) {
                if parentWord[p] != child[c] {countDiffChars++}
                p++
                c++
            }
            if countDiffChars == 1 {
                adjList[parentWord] = append(adjList[parentWord], child)
            }
        }
    }
    if !listHasEndWord {return 0}
    level := 1
    q := []string{beginWord}
    visited := map[string]struct{}{beginWord:struct{}{}}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return level}
            for _, nei := range adjList[dq] {
                if _, ok := visited[nei]; !ok {
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