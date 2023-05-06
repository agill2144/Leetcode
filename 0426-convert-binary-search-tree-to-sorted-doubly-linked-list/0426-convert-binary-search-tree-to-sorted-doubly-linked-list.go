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
    head := &Node{Val: 0}
    var prev *Node
    tail := head
    
    var dfs func(r *Node)
    dfs = func(r *Node) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Left)
        tail.Right = r
        tail = tail.Right
        if prev != nil {
            tail.Left = prev
        }
        prev = tail
        dfs(r.Right)
    }
    dfs(root)
    newHead := head.Right
    newHead.Left = tail
    tail.Right = newHead
    return newHead
}