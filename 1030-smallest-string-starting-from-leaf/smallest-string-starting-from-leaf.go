/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: dfs without a global var
*/
// func smallestFromLeaf(root *TreeNode) string {
//     var dfs func(r *TreeNode, path []int) string
//     dfs = func(r *TreeNode, path []int) string {
//         // base
//         // return the last possible char so it fails in smallest comparison everytime
//         if r == nil {return "~"}

//         // logic
//         path = append(path, r.Val)
//         left := dfs(r.Left, path)
//         // when we are at a leaf,
//         // reverse the path and return the reversed path in string
//         if r.Left == nil && r.Right == nil {
//             return reverse(path)
//         }
//         right := dfs(r.Right, path)
//         path = path[:len(path)-1]
//         if strings.Compare(left, right) == -1 {return left}
//         return right
//     }
//     return dfs(root, nil)
// }

/*
    approach: dfs using global var
*/
func smallestFromLeaf(root *TreeNode) string {
    smallestStr := ""
    var dfs func(r *TreeNode, path string)
    dfs = func(r *TreeNode, path string) {
        // base
        if r == nil {return}

        // logic
        path = string(r.Val+'a') + path
        dfs(r.Left, path)
        dfs(r.Right, path)
        if r.Left == nil && r.Right == nil {
            // The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
            if smallestStr == "" || strings.Compare(path, smallestStr) == -1 {
                smallestStr = path
            }
        }
        
    }
    dfs(root, "")
    return smallestStr
}


