// stoneX is part of stoneY (in a connected component context)
// if stoneX share the same row or col with stoneY
// therefore the adjList: {stoneXIdx: stoneYIdx} 
func removeStones(stones [][]int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(stones); i++ {
        pRow, pCol := stones[i][0], stones[i][1]
        for j := 0; j < len(stones); j++ {
            if i == j {continue}
            currRow, currCol := stones[j][0], stones[j][1]
            if currRow == pRow || currCol == pCol {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
    }

    visited := map[int]bool{}
    var dfs func(node int)
    dfs = func(node int) {
        // base
        if visited[node] {return}
        // logic
        visited[node] = true
        for _, nei := range adjList[node] {
            dfs(nei)
        }
    }    
    
    connectedCompCount := 0
    for i := 0; i < len(stones); i++ {
        if !visited[i] {
            dfs(i)
            connectedCompCount++
        }
    }
    return len(stones) - connectedCompCount
}