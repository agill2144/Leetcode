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
    var dfs func(r *Node) *Node
    dfs = func(r *Node) *Node {
        // base
        if r == nil {return nil}

        // logic
        tail := r
        for r != nil {
            next := r.Next
            if r.Child != nil {
                ct := dfs(r.Child)
                ch := r.Child
                r.Next = ch
                ch.Prev = r
                ct.Next = next
                if next != nil {next.Prev = ct}
                if next == nil {next = ct}    
                r.Child = nil            
            }
            tail = r 
            r = next
        }
        return tail
    }
    dfs(root)
    return root
}