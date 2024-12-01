/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {return node}
    origToCp := map[*Node]*Node{}
    var dfs func(curr *Node)
    dfs = func(curr *Node) {
        // base
        if _, ok := origToCp[curr]; ok {return}

        // logic
        clone := &Node{Val: curr.Val}
        origToCp[curr] = clone
        for _, nei := range curr.Neighbors {
            dfs(nei)
            clone.Neighbors = append(clone.Neighbors, origToCp[nei])
        }
    }
    dfs(node)
    return origToCp[node]
}