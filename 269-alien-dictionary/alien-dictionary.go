func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    for i := 0; i < len(words); i++ {
        for k := 0; k < len(words[i]); k++ {
            adjList[words[i][k]] = []byte{}
        }
    }
    indegrees := make([]int, 26)
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        w1, w2 := 0, 0
        for w1 < len(word1) && w2 < len(word2) {
            if word1[w1] != word2[w2] {
                adjList[word1[w1]] = append(adjList[word1[w1]], word2[w2])
                w2Idx := int(word2[w2]-'a')
                indegrees[w2Idx]++
                break
            }
            w1++; w2++
        }
        if w2 == len(word2) && w1 < len(word1) {return ""}
    }
    countToBeProcessed := 0
    q := []byte{}
    for i := 0; i < len(indegrees); i++ {
        char := byte('a'+i)
        if adjList[char] != nil {
            countToBeProcessed++
            if indegrees[i] == 0 {q = append(q, char)}
        }
    }
    if len(q) == 0 {return ""}
    out := new(strings.Builder)
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        out.WriteByte(dq)
        countToBeProcessed--
        for _, nei := range adjList[dq] {
            indegrees[int(nei-'a')]--
            if indegrees[int(nei-'a')] == 0 {q = append(q, nei)}
        }
    }
    if countToBeProcessed > 0 {return ""}
    return out.String()
}