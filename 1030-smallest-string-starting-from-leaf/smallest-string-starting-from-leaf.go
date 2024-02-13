/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    smallestStr := ""
    var dfs func(r *TreeNode, path []int)
    dfs = func(r *TreeNode, path []int) {
        // base
        if r == nil {return}

        // logic
        path = append(path, r.Val)
        dfs(r.Left, path)
        dfs(r.Right, path)
        if r.Left == nil && r.Right == nil {
            rev := reverse(path)
            // The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
            if smallestStr == "" || strings.Compare(rev, smallestStr) == -1 {
                smallestStr = rev
            }
        }
        path = path[:len(path)-1]
    }
    dfs(root, nil)
    return smallestStr
}


func reverse(p []int) string {
    out := new(strings.Builder)
    right := len(p)-1
    for right >= 0 {
        out.WriteString(string(p[right]+'a'))
        right--
    }
    return out.String()
}