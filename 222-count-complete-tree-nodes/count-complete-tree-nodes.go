/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    var getH func(r *TreeNode) (int, int)
    getH = func(r *TreeNode) (int,int) {
        l := r
        lh := 1
        for l.Left != nil {
            l = l.Left
            lh++
        }
        rh := 1
        for r.Right != nil {
            r = r.Right
            rh++
        }
        return lh, rh
    }
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {
            return 0
        }

        // logic
        lh, rh := getH(r)
        if lh == rh {
            return int(math.Pow(2.0, float64(lh))) - 1
        } else {
            return 1 + dfs(r.Left) + dfs(r.Right)
        }
    }
    return dfs(root)
}
// brute force
// func countNodes(root *TreeNode) int {
//     count := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode){
//         // base
//         if r == nil {return}

//         //logic
//         count++
//         dfs(r.Left)
//         dfs(r.Right)
//     }
//     dfs(root)
//     return count
// }