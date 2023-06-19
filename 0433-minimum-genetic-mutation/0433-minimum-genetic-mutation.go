func minMutation(startGene string, endGene string, bank []string) int {
    adjList := map[string][]string{}
    foundEnd := false
    foundStart := false 
    for i := 0; i < len(bank); i++ {
        if bank[i] == endGene {foundEnd = true}
        if bank[i] == startGene{foundStart = true}
    }
    if !foundEnd {return -1}
    if !foundStart {bank = append(bank, startGene)}

    for i := 0; i < len(bank); i++ {
        parent := bank[i]
        for j := 0; j < len(bank); j++ {
            if i == j {continue}
            diff := 0
            child := bank[j]
            p,c := 0,0
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
    
    visited := map[string]struct{}{}
    q := []string{startGene}
    level := 0
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