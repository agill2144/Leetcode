func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    adjList := map[int][]int{} // {empID : [subordinates] }
    for i := 0; i < len(manager); i++ {
        emp := i
        man := manager[i]
        if man == -1 {continue} // this emp does not have a manager because they are the head
        adjList[man] = append(adjList[man], emp)
    }
    maxTime := 0
    var dfs func(node int, time int) 
    dfs = func(node int, time int)  {
        // base
        if time > maxTime {maxTime = time}
        
        // logic
        for _, nei := range adjList[node] {
            dfs(nei, time + informTime[node])
        }
    }
    
    dfs(headID,0)
    return maxTime
}

