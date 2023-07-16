/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    adjList := map[int][]int{}
    var dfs func(r *TreeNode, p *TreeNode)
    dfs = func(r, p *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if p != nil {adjList[r.Val] = append(adjList[r.Val], p.Val)}
        if r.Left != nil {adjList[r.Val] = append(adjList[r.Val], r.Left.Val); dfs(r.Left, r)}
        if r.Right != nil {adjList[r.Val] = append(adjList[r.Val], r.Right.Val); dfs(r.Right, r)}
    }
    dfs(root, nil)
    
    visited := map[int]struct{}{}
    visited[target.Val] = struct{}{}
    q := []int{target.Val}
    levels := 0
    out := []int{}

    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if levels == k {
                out = append(out, dq)
                qSize--
                continue
            }
            for _,nei := range adjList[dq] {
                if _, ok := visited[nei]; !ok {
                    q = append(q, nei)
                    visited[nei] = struct{}{}
                }
            }
            qSize--
        }
        levels++
    }
    return out
}