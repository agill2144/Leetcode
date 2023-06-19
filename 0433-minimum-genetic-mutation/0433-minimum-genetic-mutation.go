// identical to word-ladder
func minMutation(startGene string, endGene string, bank []string) int {
    adjList := map[string][]string{}
    foundEnd := false
    foundStart := false
    
    // o(b)
    for i := 0; i < len(bank); i++ {
        if bank[i] == endGene {foundEnd = true}
        if bank[i] == startGene{foundStart = true}
    }
    if !foundEnd {return -1}
    if !foundStart {bank = append(bank, startGene)}

    // o(b)
    for i := 0; i < len(bank); i++ {
        parent := bank[i]
        // o(b)
        for j := 0; j < len(bank); j++ {
            if i == j {continue}
            diff := 0
            child := bank[j]
            p,c := 0,0
            // o(k) where k is avg len of word
            for p < len(parent) && c < len(child) {
                if parent[p] != child[c] {
                    diff++
                }
                c++
                p++
            }
            if diff == 1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }
    
    // time so far = o(b) + o(b)xo(b)xo(k) = o(b) + o(b^2) * o(k)
    
    
    visited := map[string]struct{}{}
    q := []string{startGene}
    level := 0
    // o(b+e) ; where e is total number of edges
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endGene {return level}
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
    return -1

}