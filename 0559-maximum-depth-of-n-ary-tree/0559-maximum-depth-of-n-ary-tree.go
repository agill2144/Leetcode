/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    max := 0
    var dfs func(r *Node, depth int)
    dfs = func(r *Node, depth int) {
        // base
        if r == nil {return}
        
        // logic
        depth++
        if r.Children == nil || len(r.Children) == 0 {
            if depth > max {max = depth}
            return
        }
        for _, child := range r.Children {
            dfs(child, depth)
        }
    }
    dfs(root, 0)
    return max 
}