/*
To determine whether bomb 1 detonates bomb 2, 
we can compare the Euclidean distance between their centers and the radius of bomb 1. 
If the distance is less than or equal to the radius of bomb 1, then bomb 1 can detonate bomb 2.
Note that this relationship is not commutative: bomb 1 detonating bomb 2 does not necessarily imply the converse is also true.

given we have 2 points to compare (x1,y1) and (x2,y2)
dist = (x1-x2)^2 + (y1-y2)^2
*/
func maximumDetonation(bombs [][]int) int {
    adjList := map[int][]int{} // idx: [idx]
    for i := 0; i < len(bombs); i++ {
        parent := bombs[i]
        for j := 0; j < len(bombs); j++ {
            if i == j {continue}
            child := bombs[j]
            dist := math.Pow((float64(parent[0])-float64(child[0])), 2.0) + math.Pow((float64(parent[1])-float64(child[1])), 2)
            if dist <= float64(parent[2] * parent[2]) {
                adjList[i] = append(adjList[i], j)
            }
        }
    }
    // number of nodes in a path, and we have to find longest such path
    var dfs func(nodeIdx int, pathLen int, visited map[int]struct{})
    dfs = func(nodeIdx int, pathLen int, visited map[int]struct{}) {
        // base
        if _, ok := visited[nodeIdx]; ok {return}
        
        
        // logic
        visited[nodeIdx] = struct{}{}
        for _,nei := range adjList[nodeIdx] {
            dfs(nei, pathLen+1, visited)
        }
    }
    
    maxNodesVisited := 0
    for i := 0; i < len(bombs); i++ {
        visited := map[int]struct{}{}
        dfs(i, 1, visited)
        if len(visited) > maxNodesVisited {maxNodesVisited = len(visited)}
    }
    return maxNodesVisited
}