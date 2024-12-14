/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    copies := map[*Node]*Node{}
    var dfs func(curr *Node)
    dfs = func(curr *Node) {
        // base
        if curr == nil {return}
        if copies[curr] != nil {return}

        // logic
        copies[curr] = &Node{Val:curr.Val}
        for _, nei := range curr.Neighbors {
            dfs(nei)
            copies[curr].Neighbors = append(copies[curr].Neighbors, copies[nei])
        }
    }
    dfs(node)
    return copies[node]
}