/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// shortest path without custom weights = vanilla bfs should work

func countPairs(root *TreeNode, distance int) int {
    adjList := map[*TreeNode]map[*TreeNode]struct{}{}
    leafs := map[*TreeNode]struct{}{}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        if adjList[r] == nil {adjList[r] = map[*TreeNode]struct{}{}}
        if r.Left != nil {
            if adjList[r.Left] == nil {adjList[r.Left] = map[*TreeNode]struct{}{}}
            adjList[r][r.Left] = struct{}{}
            adjList[r.Left][r] = struct{}{}
            dfs(r.Left)

        }
        if r.Right != nil {
            if adjList[r.Right] == nil {adjList[r.Right] = map[*TreeNode]struct{}{}}
            adjList[r][r.Right] = struct{}{}
            adjList[r.Right][r] = struct{}{}        
            dfs(r.Right)

        }
        if r.Left == nil && r.Right == nil {
            leafs[r] = struct{}{}
            return
        }
    }
    dfs(root)
    count := 0
    for leaf, _ := range leafs {
        q := []*TreeNode{leaf}
        level := 0
        visited := map[*TreeNode]struct{}{}
        for len(q) != 0 {
            qSize := len(q)
            for qSize != 0 {
                dq := q[0]
                q = q[1:]
                if _, ok := leafs[dq]; ok && level <= distance && dq != leaf {
                    count++
                }
                for nei, _ := range adjList[dq] {
                    if _, ok := visited[nei]; !ok {
                        visited[nei] = struct{}{}
                        q = append(q, nei)
                    }
                }
                qSize--
            }
            level++
            if level > distance {break}
        }
    }
    return count / 2
}