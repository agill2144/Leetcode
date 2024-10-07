/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
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
    total := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        lh, rh := getH(root)
        if lh == rh {
            total += int(math.Pow(2.0,float64(lh)))-1
        } else {
            total++
            dfs(r.Left)
            dfs(r.Right)
        }
    }
    dfs(root)
    return total
}