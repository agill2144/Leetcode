/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    deepCp := map[*Node]*Node{}
    var dfs func(n *Node)
    dfs = func(n *Node) {
        // base
        if n == nil {return}
        
        // logic
        cp := &Node{Val: n.Val}
        deepCp[n] = cp
        for _, nei := range n.Neighbors {
            if val, ok := deepCp[nei]; ok {
                cp.Neighbors = append(cp.Neighbors, val)
            } else {
                dfs(nei)
                cp.Neighbors = append(cp.Neighbors, deepCp[nei])
            }
        }
    }
    dfs(node)

    return deepCp[node]
}