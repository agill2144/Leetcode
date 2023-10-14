/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if root == p || root == q {
            return root
        }
        if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        } else if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
        } else {
            return root
        }
    }
    return nil
}