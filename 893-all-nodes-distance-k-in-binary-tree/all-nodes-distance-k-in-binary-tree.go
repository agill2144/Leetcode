/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    adjList := map[*TreeNode][]*TreeNode{}
    var buildGraph func(r *TreeNode)
    buildGraph = func(r *TreeNode) {
        // base
        if r == nil {return }

        // logic
        if r.Left != nil {
            adjList[r] = append(adjList[r], r.Left)
            adjList[r.Left] = append(adjList[r.Left], r)
            buildGraph(r.Left)
        }
        if r.Right != nil {
            adjList[r] = append(adjList[r], r.Right)
            adjList[r.Right] = append(adjList[r.Right], r)
            buildGraph(r.Right)
        }
    }
    buildGraph(root)
    out := []int{}
    var dfs func(node, prev *TreeNode, depth int)
    dfs = func(node, prev *TreeNode, depth int) {
        // base
        if node == nil {return}

        // logic
        if depth == k {
            out = append(out, node.Val)
            return
        }
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            dfs(nei, node,depth+1)
        }
    }
    dfs(target, nil, 0)
    return out
}