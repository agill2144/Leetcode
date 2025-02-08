/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// naive heap tc = o(nlogk)
// without heap tc , using dq = o(n)
//
func closestKValues(root *TreeNode, target float64, k int) []int {
    if root == nil || k == 0 {return nil}
    dq := []int{}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left)
        dq = append(dq, r.Val)
        if len(dq) > k {
            if abs(float64(dq[0]) - target) < 
                abs(float64(dq[len(dq)-1])-target) {
                    dq = dq[:len(dq)-1]
                } else {
                    dq = dq[1:]
                }
        }
        dfs(r.Right)
    }
    dfs(root)
    return dq
}

func abs(x float64) float64 {
    if x < 0 {return x * -1.0}
    return x
}
