/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // morris, using threaded tree
func inorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    curr := root
    out := []int{}
    // inorder is left root right
    for curr != nil {
        if curr.Left == nil {
            // process root
            out = append(out, curr.Val)
            // go right
            curr = curr.Right
        } else {
            // connect the thread so we can come back to this node
            // from the last node of this subtree
            tmp := curr.Left
            for tmp.Right != nil && tmp.Right != curr {
                tmp = tmp.Right
            }
            if tmp.Right == nil {
                tmp.Right = curr
                // go left
                curr = curr.Left
            } else {
                // done with left subtree
                // process root
                out = append(out, curr.Val)
                // disconnect the thread
                tmp.Right = nil
                // go right
                curr = curr.Right
            }
        }
    }
    return out
}