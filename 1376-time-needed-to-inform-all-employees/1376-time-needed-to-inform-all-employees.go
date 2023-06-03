func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(manager); i++ {
        emp := i
        man := manager[emp]
        if man == -1 {continue}
        adjList[man] = append(adjList[man], emp)
    }
    
    var dfs func(node int) int
    dfs = func(node int) int {
        // base
        
        // logic
        currTime := informTime[node]
        maxTimeFromNei := 0
        for _, nei := range adjList[node] {
            if t := dfs(nei); t > maxTimeFromNei {
                maxTimeFromNei = t
            }
        }
        return currTime + maxTimeFromNei
    }
    return dfs(headID)
}