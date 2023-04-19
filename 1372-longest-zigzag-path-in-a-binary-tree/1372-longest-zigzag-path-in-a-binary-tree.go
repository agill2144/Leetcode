/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    var dfs func(r *TreeNode, goLeft bool, size int)
    ans := 0
    dfs = func(r *TreeNode, goLeft bool, size int) {
        // base
        if r == nil {return}
        
        ans = max(ans, size)
        
        
        // logic
        if goLeft {
            dfs(r.Left, false, size+1)
            dfs(r.Right, true, 1)
        } else { // supposed to go right
            dfs(r.Left, false, 1)
            dfs(r.Right, true, size+1)
        }

    }
    dfs(root, true, 0)
    dfs(root, false, 0)
    return ans
}

func max(x, y int) int {
    if x > y {return x}
    return y
}

// func longestZigZag(root *TreeNode) int {
//     q := [][]interface{}{{root, 0, 0}}
//     // 1 = right
//     // -1 = left
//     max := 0
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
//         node := dq[0].(*TreeNode)
//         size := dq[1].(int)
//         dir := dq[2].(int)
//         if size > max {max = size}
        
//         if node.Left != nil {
//             newLeftSize := 1
//             if dir == 0 || dir == 1 {
//                 newLeftSize += size
//             }
//             q = append(q, []interface{}{node.Left, newLeftSize, -1})
//         }
//         if node.Right != nil {
//             newRightSize := 1
//             if dir == 0 || dir == -1 {
//                 newRightSize += size
//             }
//             q = append(q, []interface{}{node.Right, newRightSize, 1})
//         }
//     }
    
    
//     return max
// }