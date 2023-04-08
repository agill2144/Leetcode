func longestCycle(edges []int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        if edges[i] == -1 {continue}
        adjList[i] = append(adjList[i], edges[i])
    }
    visited := make([]bool, len(edges))
    ans := -1
    
    var dfs func(n int, size int, path map[int]int)
    dfs = func(n int, size int, path map[int]int) {
        // base
        if sizeAtThatNode, ok := path[n]; ok {
            // cycle detected
            if size-sizeAtThatNode > ans {
                ans = size-sizeAtThatNode
            }
            return
        }
        if visited[n] {return}
        
        // logic
        visited[n] = true
        path[n] = size
        for _, nei := range adjList[n] {
            dfs(nei, size+1, path)
        }
        delete(path, n)
    }
    
    for i := 0; i < len(edges); i++ {
        if !visited[i]{
            dfs(i, 1 , map[int]int{})            
        }
    }
    return ans
}