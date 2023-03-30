// brute force dfs
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
    
    // EVEN tho the graph is connected, we still need to run this ON ALL NODES
    // because its possible that the pair we discover as redundant from node 1 might
    // not be the only edge ( can be but there can be edge that we discovered from other side 
    // that is at a larger idx position input )
    // if there are multiple such edges, we have to return the one that appears the last
    // lesson learned, we can run dfs on all nodes and reset visited each time 
    // to find ALL edges from all sides that maybe redundant
    // time: o(n) to form adjList + o(n) to loop * o(n) for dfs
    // total time: o(n) + o(n^2) = o(n^2)
    // space: o(n) adjList + o(n) recursion stack
    for i := 1; i < len(edges); i++ {
        visited = map[int]struct{}{}
        dfs(i, -1)
    }
    return ans
}