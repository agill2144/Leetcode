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
        if (p.Val < root.Val && q.Val > root.Val) ||
          (p.Val > root.Val && q.Val < root.Val) ||   
         (root == p || root == q) {return root}
        if p.Val < root.Val && q.Val < root.Val {root = root.Left}
        if p.Val > root.Val && q.Val > root.Val {root = root.Right}
    }
    return nil
}