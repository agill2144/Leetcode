// we are just looking for max depth with custom depth value at each depth
// custom depth = informTime
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    adjList := map[int][]int{}
    for i := 0; i < len(manager); i++ {
        p,c := manager[i],i
        if p == -1 {continue}
        adjList[p] = append(adjList[p], c)
    }
    var dfs func(node int) int
    dfs = func(node int) int {
        // base

        // logic
        t := informTime[node]
        maxChildT := 0
        for _, child := range adjList[node] {
            maxChildT = max(maxChildT, dfs(child))
        }
        return maxChildT + t
    }
    
    return dfs(headID)
}