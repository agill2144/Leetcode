/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    out := []int{}
    var dfs func(node *Node) 
    dfs = func(node *Node) {
        // base
        if node == nil {return}

        // logic
        for _, child := range node.Children {
            dfs(child)
        }
        out = append(out, node.Val)
    }
    dfs(root)
    return out
}