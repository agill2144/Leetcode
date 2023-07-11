/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
    // binary tree is an undirected graph
    // start = start node in graph
    // build adjList and start dfs/bfs from our start node ;) 
    
    adjList := map[int][]int{}
    var buildGraph func(r *TreeNode, parent *TreeNode)
    buildGraph = func(r, parent *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if parent != nil {adjList[r.Val] = append(adjList[r.Val], parent.Val)}
        if r.Left != nil {adjList[r.Val] = append(adjList[r.Val], r.Left.Val)}
        if r.Right != nil {adjList[r.Val] = append(adjList[r.Val], r.Right.Val)}
        buildGraph(r.Left, r)
        buildGraph(r.Right, r)
    }
    
    buildGraph(root, nil)
    fmt.Println(adjList)
    q := []int{start}
    visited := map[int]struct{}{}
    visited[start] = struct{}{}
    
    levels := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            
            for _, nei := range adjList[dq] {
                if _, ok := visited[nei]; !ok {
                    q = append(q, nei)
                    visited[nei] = struct{}{}
                }
            }
            qSize--
        }
        levels++
    }
    return levels-1
}