func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    type adjNode struct {
        weight float64
        node string
    }
    adjList := map[string][]*adjNode{}
    for i := 0; i < len(equations); i++ {
        a := equations[i][0]
        b := equations[i][1]
        w := values[i]
        adjList[a] = append(adjList[a], &adjNode{node: b, weight: w})
        adjList[b] = append(adjList[b], &adjNode{node: a, weight: 1.0/w})
    }
    ans := []float64{}
    
    var dfs func(node, end string, prod float64, visited map[string]struct{}) float64
    dfs = func(node, end string, prod float64, visited map[string]struct{}) float64 {
        // base
        if node == end {return prod}
        if _, ok := visited[node]; ok {return -1.0}
        
        // logic
        visited[node] = struct{}{}
        for _, nei := range adjList[node] {
            newProd := prod * nei.weight
            if val := dfs(nei.node, end, newProd, visited); val != -1.0 {delete(visited,node); return val}
        }
        delete(visited, node)
        return -1.0
    }

    v := map[string]struct{}{}
    // o(q) * o(numberOfEquations)
    for i := 0; i < len(queries); i++ {
        src := queries[i][0]
        dest := queries[i][1]
        if _, ok := adjList[src]; !ok {ans = append(ans, -1.0); continue}
        if _, ok := adjList[dest]; !ok {ans = append(ans, -1.0); continue}
        ans = append(ans, dfs(src, dest, 1.0, v))
    }
    return ans
}

