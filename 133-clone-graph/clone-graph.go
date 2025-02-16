/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {return nil}
    origToCp := map[*Node]*Node{}
    var dfs func(curr *Node) 
    dfs = func(curr *Node) {
        // base
        // if curr == nil {return}
        if origToCp[curr] != nil {return}

        // logic
        cp := &Node{Val: curr.Val}
        origToCp[curr] = cp
        for _, child := range curr.Neighbors {
            dfs(child)
            cp.Neighbors = append(cp.Neighbors, origToCp[child])
        }
    }
    dfs(node)
    return origToCp[node]
}