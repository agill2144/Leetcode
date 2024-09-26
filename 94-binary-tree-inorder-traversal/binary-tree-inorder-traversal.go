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
    for curr != nil {
        if curr.Left == nil {
            // left is nil
            // process root ( left(nil) -> root(we are here) -> right(next step) )
            out = append(out, curr.Val)
            curr = curr.Right
        } else {
            tmp := curr.Left
            for tmp.Right != nil && tmp.Right != curr {
                tmp = tmp.Right
            }
            if tmp.Right == nil {
                // create the thread to curr
                tmp.Right = curr
                // go left
                curr = curr.Left
            } else {
                // tmp Right is curr, i.e this is a thread we created
                // i.e from left, we came back to parent / root ( left(done) -> root(we are here) -> right(next step) )
                
                // disconnect tmp -> curr
                tmp.Right = nil
                // process curr ( i.e root )
                out = append(out, curr.Val)
                // go right
                curr = curr.Right
            }
        }
    }
    return out
}