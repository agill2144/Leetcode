/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {return true}
    if root1 == nil || root2 == nil {return false}
    
    r1Leafs := []int{}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {
            return
        }
        
        // logic
        dfs(r.Left)
        if r.Left == nil && r.Right == nil {
            r1Leafs = append(r1Leafs, r.Val)
        }
        dfs(r.Right)
    }
    dfs(root1)
    r1 := make([]int, len(r1Leafs))
    copy(r1, r1Leafs)
    r1Leafs = []int{}
    dfs(root2)
    
    if len(r1) != len(r1Leafs) {return false}
    for i := 0; i < len(r1); i++ {
        if r1[i] != r1Leafs[i] {return false}
    }
    return true
}