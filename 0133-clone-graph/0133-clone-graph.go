/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    deepCp := map[*Node]*Node{}
    var dfs func(root *Node)
    dfs = func(root *Node) {
        // base
        if root == nil {return}
        if _, ok := deepCp[root]; ok {return}
        
        // logic
        clone := &Node{Val: root.Val}
        deepCp[root] = clone
        for _, nei := range root.Neighbors {
            dfs(nei)
            clone.Neighbors = append(clone.Neighbors, deepCp[nei])
        }
    }
    dfs(node)
    return deepCp[node]
}