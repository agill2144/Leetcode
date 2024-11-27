/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList(root *Node) *Node {
    if root == nil {return nil}
    dummy := &Node{Val: 0}
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