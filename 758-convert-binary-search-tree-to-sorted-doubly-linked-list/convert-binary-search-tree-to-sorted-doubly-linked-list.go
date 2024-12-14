/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */
// use tail as prev ptr
// after we have appended a node to tail
// tail is prev val
// meaning tail -> currNode
// to make it doubly-LL
// curr needs to point back at prev node
// TAIL is PREV NODE!!
// tail <-> currNode
// now move tail ptr to newewly appended node ( because now this node is prev )
func treeToDoublyList(root *Node) *Node {
    if root == nil {return nil}
    dummy := &Node{Val:0}
    tail := dummy
    var dfs func(r *Node)
    dfs = func(r *Node) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left)
        tail.Right = r
        r.Left = tail
        tail = tail.Right
        dfs(r.Right)
    }
    dfs(root)
    head := dummy.Right
    head.Left = tail
    tail.Right = head
    return head
}