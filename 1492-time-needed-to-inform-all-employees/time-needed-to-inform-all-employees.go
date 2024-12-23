// we are just looking for max depth with custom depth value at each depth
// custom depth = informTime
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(manager); i++ {
        p,c := manager[i],i
        if p == -1 {continue}
        adjList[p] = append(adjList[p], c)
    }
    maxTime := 0
    var dfs func(node, t int)
    dfs = func(node, t int) {
        // base
        
        maxTime = max(maxTime, t)
        // logic
        for _, child := range adjList[node] {
            dfs(child, t+informTime[node])
        }
    }
    dfs(headID, 0)
    return maxTime
}