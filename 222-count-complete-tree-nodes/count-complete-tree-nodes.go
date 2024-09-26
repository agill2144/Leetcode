/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // number of nodes in a complete full binary tree = 2^h - 1
 // complete full = leftHeight == rightHeight
func countNodes(root *TreeNode) int {
    if root == nil {return 0}
    getH := func(r *TreeNode) (int, int) {
        lh := 0
        l := r
        for l != nil {
            lh++
            l = l.Left
        }
        rh := 0
        for r != nil {
            rh++
            r = r.Right
        }
        return lh, rh
    }
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        lh, rh := getH(r)
        if lh == rh {
            return int(math.Pow(2.0, float64(lh)))-1
        } else {
            return 1 + dfs(r.Left) + dfs(r.Right)
        }
    }
    return dfs(root)
}