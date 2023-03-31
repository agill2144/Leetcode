func eventualSafeNodes(graph [][]int) []int {
    var ans []int
    // [node]0:white/1:grey/2:black
    visited := make(map[int]int, len(graph))
    // check each each starting point if it has a cycle
    for start := 0; start < len(graph); start++ {
        if !hasCycle(graph, start, visited) {
            ans = append(ans, start)
        }
    }
    return ans
}

// hasCycle dfs will check graph for cycles from this node.
func hasCycle(graph [][]int, node int, visited map[int]int) bool {
    // mark node as visiting
    visited[node] = 1 // begin visiting this node
    // check children for cycle
    for _, child := range graph[node] {
        // If this node path has been checked, skip it with answer.
        if visited[child] == 1 {
            return true 
        }
        // Only visit a path if it hasn't been visited already.
        if visited[child] == 0 && hasCycle(graph, child, visited) {
            return true
        }
    }
    // mark node as visited
    visited[node] = 2
    return false
}