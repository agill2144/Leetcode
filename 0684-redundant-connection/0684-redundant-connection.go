func findRedundantConnection(edges [][]int) []int {
    set := map[[2]int]int{}
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        set[[2]int{src,dest}] = i
        adjList[src] = append(adjList[src], dest)
        adjList[dest] = append(adjList[dest], src)        
    }
    visited := map[int]struct{}{}
    var ans []int
    ansIdx := -1
    var dfs func(node, prev int)
    dfs = func(node, prev int)  {
        // base
        if _, ok := visited[node]; ok {
            pair := [2]int{node,prev}
            if prev < node { pair = [2]int{prev, node}}
            if idx, found := set[pair]; found {
                if ans == nil || idx > ansIdx {
                    ans = []int{pair[0], pair[1]}
                    ansIdx = idx
                }
            }
            return
        }
        
        // logic
        visited[node] = struct{}{}
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node)
        }
    }
    for i := 1; i < len(edges); i++ {
        visited = map[int]struct{}{}
        dfs(i, -1)        
    }
    return ans
}