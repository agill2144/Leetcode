
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
    // root is a node with 0 incoming edges and should have left or right or both childs
    root := -1
    for i := 0; i < n; i++ {
        incomingEdges := indegrees[i]
        hasLeft := leftChild[i] != -1
        hasRight := rightChild[i] != -1
        if incomingEdges == 0 && (hasLeft || hasRight) {
            // root was discovered but we found another root,
            // if we find multiple such roots, return false right away
            // a valid binary tree only has 1 root
            if root != -1 {return false}
            root = i
        }
    }
    if root == -1 {return false}
    visited := make([]bool, n)
    // make sure we visited all of them
    // if there were disconnected components
    // we would not be able to visit all of them from single root
    // therefore not a valid binary tree
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