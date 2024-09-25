/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    out := []int{}
    curr := root
    // left -> root -> right
    for curr != nil {
        if curr.Left == nil {
            // when there is no left child
            // this is root node in inorder, process it
            // then go right
            out = append(out, curr.Val)
            curr = curr.Right
        } else {
            // create a thread from last right child back to curr
            tmp := curr.Left
            for tmp.Right != nil && tmp.Right != curr {
                tmp = tmp.Right
            }
            if tmp.Right == nil {
                // create the thread from right most to curr
                tmp.Right = curr
                // we have created a thread to come back to root
                // now go left
                curr = curr.Left
            } else {
                // left subtree is done
                // disconnect the thread
                tmp.Right = nil
                // this is root node in "inorder" , process it
                out = append(out, curr.Val)
                // root is done, now go right
                curr = curr.Right
            }
        }
    }
    return out
}