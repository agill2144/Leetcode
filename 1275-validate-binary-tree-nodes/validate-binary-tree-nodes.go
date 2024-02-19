
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    if n <= 1 {return true}
    adjList := map[int][][]int{}
    indegrees := make([]int, n)
    for i := 0; i < n; i++ {
        adjList[i] = append(adjList[i], []int{leftChild[i], rightChild[i]})
        if leftChild[i] != -1 {
            indegrees[leftChild[i]]++
        }
        if rightChild[i] != -1 {
            indegrees[rightChild[i]]++
        }
    }
    root := -1
    for i := 0; i < n; i++ {
        incomingEdges := indegrees[i]
        hasLeft := leftChild[i] != -1
        hasRight := rightChild[i] != -1
        if incomingEdges == 0 && (hasLeft || hasRight) {
            if root != -1 {return false}
            root = i
        }
    }
    if root == -1 {return false}
    fmt.Println(root)
    visited := make([]bool, n)
    visitedCount := 0
    var dfs func(node int) bool
    dfs = func(node int) bool {
        // base
        if visited[node] {return false}
        
        // logic
        visited[node] = true
        visitedCount++
        for _, nodes := range adjList[node] {
            for _, n := range nodes {
                if n == -1 {continue}
                if !dfs(n){return false}
            }
        }
        return true
    }
    return dfs(root) && visitedCount == n
}