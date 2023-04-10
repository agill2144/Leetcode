func minReorder(n int, connections [][]int) int {
    adjList := map[int][]int{}
    set := map[[2]int]struct{}{}
    for i := 0; i < len(connections); i++ {
        src := connections[i][0]
        dest := connections[i][1]
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)      
        set[[2]int{src,dest}] = struct{}{}
    }
    visited := make([]bool, n)
    count := 0
    var dfs func(node, prev int)
    dfs = func(node, prev int) {
        // base
        if visited[node] {return}
                
        // logic
        visited[node] = true
        if _, ok := set[[2]int{node,prev}]; !ok && prev != -1 {count++}
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node)
        }
    }
    dfs(0,-1)
    return count
}