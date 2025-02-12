/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    if root == nil || target == nil {return nil}
    targetFound := false
    adjList := map[*TreeNode][]*TreeNode{}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        if r == target {targetFound = true}
        if _, ok := adjList[r]; !ok {
            adjList[r] = []*TreeNode{}
        }
        if r.Left != nil {
            adjList[r] = append(adjList[r], r.Left)
            adjList[r.Left] = append(adjList[r.Left], r)
            dfs(r.Left)
        }
        if r.Right != nil {
            adjList[r] = append(adjList[r], r.Right)
            adjList[r.Right] = append(adjList[r.Right], r)
            dfs(r.Right)
        }
    }
    dfs(root)
    if !targetFound {return nil}
    q := []*TreeNode{target}
    dist := 0
    out := []int{}
    visited := map[*TreeNode]bool{target:true}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dist == k {
                out = append(out, dq.Val)
                qSize--
                continue
            } else {
                for _, child := range adjList[dq] {
                    if visited[child] {continue}
                    q = append(q, child)
                    visited[child] = true
                }
            }
            qSize--
        }
        dist++
    }
    return out
}