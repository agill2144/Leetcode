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
    out := []int{}
    var dfs2 func(node, prev, dist int)
    dfs2 = func(node, prev, dist int) {
        // base
        if dist >= k {
            if dist == k {out = append(out, node)}
            return
        }
        
        // logic
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs2(nei, node, dist+1)
        }
    }
    dfs2(target.Val, -1, 0)
    return out
    
}