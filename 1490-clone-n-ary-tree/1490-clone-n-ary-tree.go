/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func cloneTree(root *Node) *Node {
    deepCp := map[*Node]*Node{}
    var dfs func(r *Node)
    dfs = func(r *Node) {
        // base
        if r == nil {return}
        if _, ok := deepCp[r]; ok {return}
        
        // logic
        clone := &Node{Val: r.Val}
        deepCp[r] = clone
        for _, child := range r.Children{
            dfs(child)
            clone.Children = append(clone.Children, deepCp[child])
        }
    }
    dfs(root)
    return deepCp[root]
}