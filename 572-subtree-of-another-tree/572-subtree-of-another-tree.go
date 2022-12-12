/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    // same := isSameTree(root, subRoot)
    // if same {
    //     return true
    // } else {
        // left := isSubtree(root.Left, subRoot)
        // if left {
        //    return true 
        // }
        // return isSubtree(root.Right, subRoot) 
        // the above is just a left || right -- so 1 liner
        return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
    // }
    
}

func isSameTree(a, b *TreeNode) bool {
    if a == nil && b == nil {return true}
    if (a == nil && b != nil) || (a != nil && b == nil) {return false}
    return a.Val == b.Val && isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}