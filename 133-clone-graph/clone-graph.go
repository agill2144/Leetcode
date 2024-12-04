/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {return nil}
    copies := map[*Node]*Node{}
    var dfs func(curr *Node)
    dfs = func(curr *Node) {
        // base
        if copies[curr] != nil {return}

        // logic
        cp := &Node{Val: curr.Val}
        copies[curr] = cp
        for _, nei := range curr.Neighbors {
            dfs(nei)
            cp.Neighbors = append(cp.Neighbors, copies[nei])
        }
    }
    dfs(node)
    return copies[node]
}