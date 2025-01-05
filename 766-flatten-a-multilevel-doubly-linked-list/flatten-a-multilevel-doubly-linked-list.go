/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    if root == nil {return root}
    var dfs func(r *Node) *Node
    dfs = func(r *Node) *Node {
        // base
        if r == nil {return nil}
        // logic
        var prev *Node
        for r != nil {
            next := r.Next
            if r.Child != nil {
                ct := dfs(r.Child)
                r.Next = r.Child
                r.Child.Prev = r
                r.Child = nil
                ct.Next = next
                if next != nil {next.Prev = ct}
                if next == nil {next=ct}
            }
            prev = r
            r = next
        }
        return prev
    }
    dfs(root)
    return root
}