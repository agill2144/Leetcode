/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
    adjList := map[int][]int{}
    var dfs func(r *TreeNode, parent *TreeNode)
    dfs = func(r, parent *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if r.Left != nil {adjList[r.Val] = append(adjList[r.Val], r.Left.Val)}
        if r.Right != nil {adjList[r.Val] = append(adjList[r.Val], r.Right.Val)}
        if parent != nil {adjList[r.Val] = append(adjList[r.Val], parent.Val)}
        dfs(r.Left, r)
        dfs(r.Right, r)
    }
    
    // build adjList
    dfs(root, nil)
    // perform bfs
    visited := map[int]struct{}{}
    visited[start] = struct{}{}
    q := []int{start}
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            for _, nei := range adjList[dq] {
                if _, ok := visited[nei]; !ok {
                    visited[nei] = struct{}{}
                    q = append(q, nei)
                }
            }
            qSize--
        }
        level++
    }
    return level-1
}