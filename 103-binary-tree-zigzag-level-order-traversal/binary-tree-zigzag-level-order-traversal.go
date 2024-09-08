/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach:
    - level order using dfs or bfs
    - then reverse each odd level at the end
    time = o(n) * o(n/2 * n) - because prepending happens at every odd level and prepending in golang takes o(n) time 
    space = o(h) or o(n) for recursive stack incase of a skewed tree
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
    out := [][]int{}
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        // using level as a idx
        // if len(out) == level(idx)
        // it means idx/level does not exist
        // therefore add a new empty list 
        // so that idx exists 
        if len(out) == level {
            out = append(out, []int{})
        }
        if level % 2 != 0 {
            // odd level = prepend
            out[level] = append([]int{r.Val}, out[level]...)
        } else {
            out[level] = append(out[level], r.Val) 
        }
        dfs(r.Left, level+1)
        dfs(r.Right,level+1)
    }
    dfs(root, 0)
    return out
}

 /*
    2 stack approach
    - draw it out and simulate it 
    time = o(n)
    space = o(2n)
 */
// func zigzagLevelOrder(root *TreeNode) [][]int {
//     if root == nil {return nil}
//     out := [][]int{}
//     even := []*TreeNode{}
//     odd := []*TreeNode{root}
//     for len(odd) != 0 || len(even) != 0 {
//         // process odd first
//         level := []int{}
//         for len(odd) != 0 {
//             top := odd[len(odd)-1]
//             odd = odd[:len(odd)-1]
//             level = append(level, top.Val)
//             if top.Left != nil {even = append(even, top.Left)}
//             if top.Right != nil {even = append(even, top.Right)}
//         }
//         if len(level) > 0 {
//             out = append(out, level)
//         }
//         level = []int{}
//         for len(even) != 0 {
//             top := even[len(even)-1]
//             even = even[:len(even)-1]
//             level = append(level, top.Val)
//             if top.Right != nil {odd = append(odd, top.Right)}
//             if top.Left != nil {odd = append(odd, top.Left)}
//         }
//         if len(level) > 0 {
//             out = append(out, level)
//         }
//     }
//     return out
// }