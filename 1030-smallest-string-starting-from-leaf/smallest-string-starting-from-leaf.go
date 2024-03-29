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
func smallestFromLeaf(root *TreeNode) string {
    var dfs func(r *TreeNode, path []int) string
    dfs = func(r *TreeNode, path []int) string {
        // base
        // return the last possible char so it fails in smallest comparison everytime
        if r == nil {return "~"}

        // logic
        path = append(path, r.Val)
        left := dfs(r.Left, path)
        // when we are at a leaf,
        // reverse the path and return the reversed path in string
        if r.Left == nil && r.Right == nil {
            return reverse(path)
        }
        right := dfs(r.Right, path)
        path = path[:len(path)-1]
        if strings.Compare(left, right) == -1 {return left}
        return right
    }
    return dfs(root, nil)
}

/*
    approach: dfs using global var
*/
// func smallestFromLeaf(root *TreeNode) string {
//     smallestStr := ""
//     var dfs func(r *TreeNode, path []int)
//     dfs = func(r *TreeNode, path []int) {
//         // base
//         if r == nil {return}

//         // logic
//         path = append(path, r.Val)
//         dfs(r.Left, path)
//         dfs(r.Right, path)
//         if r.Left == nil && r.Right == nil {
//             rev := reverse(path)
//             // The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//             if smallestStr == "" || strings.Compare(rev, smallestStr) == -1 {
//                 smallestStr = rev
//             }
//         }
//         path = path[:len(path)-1]
//     }
//     dfs(root, nil)
//     return smallestStr
// }


func reverse(p []int) string {
    out := new(strings.Builder)
    right := len(p)-1
    for right >= 0 {
        out.WriteString(string(p[right]+'a'))
        right--
    }
    return out.String()
}