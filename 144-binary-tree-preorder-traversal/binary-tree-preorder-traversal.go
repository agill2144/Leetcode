/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    out := []int{}
    curr := root
    for curr != nil {
        if curr.Left == nil {
            // process root
            out = append(out, curr.Val)
            // go right
            curr = curr.Right
        } else {
            // create thread from left to root
            tmp := curr.Left
            for tmp.Right != nil && tmp.Right != curr {
                tmp = tmp.Right
            }
            if tmp.Right == nil {
                // process root
                out = append(out, curr.Val)
                // create thread from left to root
                tmp.Right = curr
                // go left
                curr = curr.Left
            } else {
                // left subtree of curr is done
                // disconnect the thread
                tmp.Right = nil
                // go right
                curr = curr.Right
            }
        }
    }
    return out
}