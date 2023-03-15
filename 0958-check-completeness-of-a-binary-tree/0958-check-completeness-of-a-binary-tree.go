/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// insert all nil nodes blindly, and as soon as we run into a nil node, flag it and do not add its child
// if we run into a regular node and we have already seen a nil node, then this cannot be a complete tree
// because a regular node should have came before a nil node if that regular node was on the farthest left side possible.
// This works because in bfs we process left to right. that is left first.
// if while processing from left to right, we encounter a nil node, flag it ( nothing to the right of this should be a regular node )
func isCompleteTree(root *TreeNode) bool {
    if root == nil {return true}

    nullSeen := false
    q := []*TreeNode{root}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq != nil && nullSeen {return false}
        if dq == nil {nullSeen = true; continue}
        
        q = append(q, dq.Left)
        q = append(q, dq.Right)
    }
    return true
}