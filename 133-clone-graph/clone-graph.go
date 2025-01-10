/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    ogToCp := map[*Node]*Node{}
    var dfs func(curr *Node)
    dfs = func(curr *Node) {
        // base
        if curr == nil {return}
        if ogToCp[curr] != nil {return}
        // logic
        ogToCp[curr] = &Node{Val:curr.Val, Neighbors: []*Node{}}
        for _, nei := range curr.Neighbors {
            dfs(nei)
            ogToCp[curr].Neighbors = append(ogToCp[curr].Neighbors, ogToCp[nei])
        }
        
    }
    dfs(node)
    return ogToCp[node]
}