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
        curr := r
        tail := r
        for curr != nil {
            next := curr.Next
            if curr.Child != nil {
                ct := dfs(curr.Child)
                curr.Next = curr.Child
                curr.Child.Prev = curr
                ct.Next = next
                // 2 cases are possible
                // next is nil, and next is not nil
                // if next is nil, meaning our new tail should be ct (childTail)
                // if next is nil, we want to return ct in this case
                // if next is not nil, connect your ct, and go to next node , keep looking for tail
                if next != nil {
                    next.Prev = ct
                } else {
                    next = ct
                }
                curr.Child = nil
            }
            tail = curr
            curr = next
        }
        return tail
    }
    dfs(root)
    return root
}